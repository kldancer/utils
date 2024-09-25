package concurrency

import (
	"fmt"
	"sync"

	"k8s.io/apimachinery/pkg/util/rand"
)

func chan1() {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {

		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {

		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 100) {
		fmt.Printf("%v ", num)
	}

}

func chan2() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {

		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {

		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Printf("%v ", num)
	}

}

func oddAndEven() {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	n := 10

	go func(n int) {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-c1
			fmt.Println("偶数：", i)
			c2 <- true
		}
	}(n)

	go func(n int) {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-c2
			fmt.Println("奇数：", i)
			c1 <- true
		}
	}(n)

	// 启动第一个goroutine，因为从1开始，所以先启动打印奇数的协程
	c2 <- true

	wg.Wait()
}
