package others

import (
	"fmt"
	"sync"
)

/*

设计三个协程 Goroutine 1， Goroutine 2， Goroutine 3 交替打印1~100
OutPut：
Goroutine 1: 1
Goroutine 2: 2
Goroutine 3: 3
Goroutine 1: 4
···
Goroutine 1: 100

*/

func multiChanPrint() {
	var i int
	x := make(chan int)
	y := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go func() {
		for {
			<-c1
			i++
			_, ok := <-y
			if !ok {
				return
			}
			fmt.Printf("Goroutine 1: %d\n", i)
			x <- i
			c2 <- struct{}{}
		}
	}()
	go func() {
		for {
			<-c2
			i++
			_, ok := <-y
			if !ok {
				return
			}
			fmt.Printf("Goroutine 2: %d\n", i)
			x <- i
			c3 <- struct{}{}
		}
	}()
	go func() {
		for {
			<-c3
			i++
			_, ok := <-y
			if !ok {
				return
			}
			fmt.Printf("Goroutine 3: %d\n", i)
			x <- i
			c1 <- struct{}{}
		}
	}()
	c1 <- struct{}{}
	y <- struct{}{}
	for k := range x {
		if k >= 100 {
			close(y)
			break
		}
		y <- struct{}{}
	}
}

func multiRoutinePrint1() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	wg.Add(3)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 3 {
			<-ch1
			fmt.Printf("GoRoutine1: %d\n", i)
			if i != 100 {
				ch2 <- struct{}{}
			}
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 3 {
			<-ch2
			fmt.Printf("GoRoutine2: %d\n", i)
			ch3 <- struct{}{}
		}
	}()

	// Goroutine 3
	go func() {
		defer wg.Done()
		for i := 3; i <= 100; i += 3 {
			<-ch3
			fmt.Printf("GoRoutine3: %d\n", i)
			ch1 <- struct{}{}
		}
	}()

	// 启动第一个goroutine
	ch1 <- struct{}{}

	wg.Wait()
}
