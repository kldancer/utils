package library

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

func b1() {
	b := bytebufferpool.Get()
	b.WriteString("hello")
	b.WriteByte(',')
	b.WriteString(" world!")

	fmt.Println(b.String())

	bytebufferpool.Put(b)
}
