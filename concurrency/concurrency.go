package concurrency

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func Concurrency() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) { //1
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) //2
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw, "%d\t%v\t%v\n", count,
			test(count, &m, m.RLocker()), test(count, &m, &m),
		)
	}

}

func Cond() {
	c := sync.NewCond(&sync.Mutex{})    //1
	queue := make([]interface{}, 0, 10) //2

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        //8
		queue = queue[1:] //9
		fmt.Println("Removed from queue")
		c.L.Unlock() //10
		c.Signal()   //11
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            //3
		for len(queue) == 2 { //4
			c.Wait() //5
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) //6
		c.L.Unlock()                        //7
	}
}

func Loop() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			fmt.Println("working...")
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
