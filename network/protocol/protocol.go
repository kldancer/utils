package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net"
	"sync"
)

// 设计思路：
// 1. 定义一个 connWriter结构体，作为io.WriteCloser 的实现，它在内部维护缓存用来存储发送数据并在 Close 时，根据缓存情况将数据一次性发送，然后清理缓存。
// 2. 传输数据时。先发送数据的长度，再发送数据内容。接收数据时，也是先接收数据长度，再根据数据长度设置byte数组长度来接收数据内容。
// 3. 传输数据流程：a) conn.Send(key)：发送key数据长度 -> 发送key数据内容. b) 通过writer.Write([]byte)：存储发送数据到writer buffer缓存. c) writer.Close触发后，发送数据长度-> 发送数据内容
// 4. 接收数据流程：a）conn.Receive()：读取key长度 -> 读取key内容 -> 读取数据长度 -> 读取数据内容

// Conn 是你需要实现的一种连接类型，它支持下面描述的若干接口；
// 为了实现这些接口，你需要设计一个基于 TCP 的简单协议；

type Conn struct {
	conn       net.Conn
	mu         sync.Mutex
	name       string
	connWriter *connWriter
}

// 用于标识传输的 writer
type connWriter struct {
	conn   *Conn
	buffer *bytes.Buffer
	key    string
}

func (w *connWriter) Write(p []byte) (n int, err error) {
	n, err = w.buffer.Write(p)
	fmt.Printf("[%s]-[Send] key: %s Writing %d bytes to buffer \n", w.conn.name, w.key, n)
	return n, err
}

func (w *connWriter) Close() error {
	w.conn.mu.Lock()
	defer w.conn.mu.Unlock()

	// 写入数据长度和数据内容
	length := uint32(w.buffer.Len())
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, length)

	if length > 0 {
		fmt.Printf("[%s]-[Send] key: %s Closing writer, Data total length: %d bytes\n", w.conn.name, w.key, length)

		// 写入长度
		_, err := w.conn.conn.Write(header)
		if err != nil {
			return err
		}

		// 写入内容
		_, err = w.conn.conn.Write(w.buffer.Bytes())
		if err != nil {
			return err
		}

		fmt.Printf("[%s]-[Send] key: %s Finished sending data \n", w.conn.name, w.key)
		fmt.Printf("[%s]-[Send] key: %s reset writer buffer \n", w.conn.name, w.key)
		w.buffer.Reset()
	}
	return nil
}

// Send 传入一个 key 表示发送者将要传输的数据对应的标识；
// 返回 writer 可供发送者分多次写入大量该 key 对应的数据；
// 当发送者已将该 key 对应的所有数据写入后，调用 writer.Close 告知接收者：该 key 的数据已经完全写入；
func (conn *Conn) Send(key string) (io.WriteCloser, error) {
	conn.mu.Lock()
	defer conn.mu.Unlock()

	// 发送 key 长度和内容
	keyLength := uint32(len(key))
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, keyLength)
	_, err := conn.conn.Write(header)
	if err != nil {
		return nil, err
	}

	_, err = conn.conn.Write([]byte(key))
	if err != nil {
		return nil, err
	}

	fmt.Printf("[%s]-[Send] Key: %s, Key length: %d bytes\n", conn.name, key, keyLength)

	// 返回一个用于写入数据的 writer
	cw := &connWriter{
		conn:   conn,
		buffer: new(bytes.Buffer),
		key:    key,
	}

	conn.connWriter = cw
	return cw, nil
}

// Receive 返回一个 key 表示接收者将要接收到的数据对应的标识；
// 返回的 reader 可供接收者多次读取该 key 对应的数据；
// 当 reader 返回 io.EOF 错误时，表示接收者已经完整接收该 key 对应的数据；
func (conn *Conn) Receive() (string, io.Reader, error) {
	conn.mu.Lock()
	defer conn.mu.Unlock()

	// 读取 key 长度和 key
	header := make([]byte, 4)
	_, err := io.ReadFull(conn.conn, header)
	if err != nil {
		return "", nil, err
	}
	keyLength := binary.BigEndian.Uint32(header)

	keyBytes := make([]byte, keyLength)
	_, err = io.ReadFull(conn.conn, keyBytes)
	if err != nil {
		return "", nil, err
	}
	key := string(keyBytes)

	// 读取数据长度
	_, err = io.ReadFull(conn.conn, header)
	if err != nil {
		return "", nil, err
	}
	dataLength := binary.BigEndian.Uint32(header)

	fmt.Printf("[%s]-[Receive] key: %s, Key length: %d bytes, Data length: %d bytes\n", conn.name, key, keyLength, dataLength)

	// 读取数据内容
	data := make([]byte, dataLength)
	_, err = io.ReadFull(conn.conn, data)
	if err != nil {
		return "", nil, err
	}

	fmt.Printf("[%s]-[Receive] key %s: Finished receiving data\n", conn.name, key)
	return key, bytes.NewReader(data), nil
}

// Close 关闭你实现的连接对象及其底层的 TCP 连接
func (conn *Conn) Close() {
	fmt.Printf("[%s]-[Close] Closing connection\n", conn.name)
	conn.connWriter.Close()
	conn.conn.Close()
}

// NewConn 从一个 TCP 连接得到一个你实现的连接对象
func NewConn(conn net.Conn, name string) *Conn {
	return &Conn{conn: conn, name: name}
}

// 除了上面规定的接口，你还可以自行定义新的类型，变量和函数以满足实现需求

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 连接到测试服务器，获得一个你实现的连接对象
func dial(serverAddr, name string) *Conn {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	return NewConn(conn, name)
}

// 启动测试服务器
func startServer(handle func(*Conn)) net.Listener {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("[WARNING] ln.Accept", err)
				return
			}
			go handle(NewConn(conn, "server"))
		}
	}()
	return ln
}

// 简单断言
func assertEqual[T comparable](actual T, expected T) {
	if actual != expected {
		panic(fmt.Sprintf("actual:%v expected:%v\n", actual, expected))
	}
}

// 简单 case：单连接，双向传输少量数据
func testCase0() {
	const (
		key  = "Bible"
		data = `Then I heard the voice of the Lord saying, “Whom shall I send? And who will go for us?”
And I said, “Here am I. Send me!”
Isaiah 6:8`
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		_key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		assertEqual(_key, key)
		dataB, err := io.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		assertEqual(string(dataB), data)

		// 服务端向客户端进行传输
		writer, err := conn.Send(key)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write([]byte(data))
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(n)
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String(), "client")
	// 客户端向服务端传输
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	n, err := writer.Write([]byte(data))
	if n != len(data) {
		panic(n)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	// 客户端等待服务端传输
	_key, reader, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	assertEqual(_key, key)
	dataB, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	assertEqual(string(dataB), data)
	conn.Close()
}

// 生成一个随机 key
func newRandomKey() string {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buf)
}

// 读取随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func readRandomData(reader io.Reader, hash hash.Hash) (checksum string) {
	hash.Reset()
	var buf = make([]byte, 23<<20) //调用者读取时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf[:n])
		if err != nil {
			panic(err)
		}
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 写入随机数据，并返回随机数据的校验和：用于验证数据是否完整传输
func writeRandomData(writer io.Writer, hash hash.Hash) (checksum string) {
	hash.Reset()
	const (
		dataSize = 500 << 20 //一个 key 对应 500MB 随机二进制数据，dataSize 也可以是其他值，你的实现中不可假定 dataSize 为固定值
		bufSize  = 1 << 20   //调用者写入时的 buf 大小不是固定的，你的实现中不可假定 buf 为固定值
	)
	var (
		buf  = make([]byte, bufSize)
		size = 0
	)
	for i := 0; i < dataSize/bufSize; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			panic(err)
		}
		_, err = hash.Write(buf)
		if err != nil {
			panic(err)
		}
		n, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
		size += n
	}
	if size != dataSize {
		panic(size)
	}
	checksum = hex.EncodeToString(hash.Sum(nil))
	return checksum
}

// 复杂 case：多连接，双向传输，大量数据，多个不同的 key
func testCase1() {
	var (
		mapKeyToChecksum = map[string]string{}
		lock             sync.Mutex
	)
	ln := startServer(func(conn *Conn) {
		// 服务端等待客户端进行传输
		key, reader, err := conn.Receive()
		if err != nil {
			panic(err)
		}
		var (
			h         = sha256.New()
			_checksum = readRandomData(reader, h)
		)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)

		// 服务端向客户端连续进行 2 次传输
		for _, key := range []string{newRandomKey(), newRandomKey()} {
			writer, err := conn.Send(key)
			if err != nil {
				panic(err)
			}
			checksum := writeRandomData(writer, h)
			lock.Lock()
			mapKeyToChecksum[key] = checksum
			lock.Unlock()
			err = writer.Close() //表明该 key 的所有数据已传输完毕
			if err != nil {
				panic(err)
			}
		}
		conn.Close()
	})
	//goland:noinspection GoUnhandledErrorResult
	defer ln.Close()

	conn := dial(ln.Addr().String(), "client")
	// 客户端向服务端传输
	var (
		key = newRandomKey()
		h   = sha256.New()
	)
	writer, err := conn.Send(key)
	if err != nil {
		panic(err)
	}
	checksum := writeRandomData(writer, h)
	lock.Lock()
	mapKeyToChecksum[key] = checksum
	lock.Unlock()
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 客户端等待服务端的多次传输
	keyCount := 0
	for {
		key, reader, err := conn.Receive()
		if err == io.EOF {
			// 服务端所有的数据均传输完毕，关闭连接
			break
		}
		if err != nil {
			panic(err)
		}
		_checksum := readRandomData(reader, h)
		lock.Lock()
		checksum, keyExist := mapKeyToChecksum[key]
		lock.Unlock()
		if !keyExist {
			panic(fmt.Sprintln(key, "not exist"))
		}
		assertEqual(_checksum, checksum)
		keyCount++
	}
	assertEqual(keyCount, 2)
	conn.Close()
}

func main() {
	testCase0()
	//testCase1()
}
