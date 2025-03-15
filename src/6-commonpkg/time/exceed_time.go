package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan struct{}, 1)
	go func() {
		LongTimeWork(cancel) //这是我们要控制超时的函数
	}()

	select { //下面的case只执行最早到来的那一个
	case <-time.After(20 * time.Second): //timeout到来
		cancel <- struct{}{}
		fmt.Printf("%s main timeout\n", time.Now())
	}
}

func LongTimeWork(cancel chan struct{}) {
	for {
		select {
		case <-cancel:
			fmt.Printf("%s LongTimeWork cancel\n", time.Now())
			return
		default:
			fmt.Printf("%s Working\n", time.Now())
			time.Sleep(3 * time.Second)
		}
	}
}

// todo 想在执行f()的同时，进行超时判断，目前函数还存在问题
func exceedTime(f func(), d time.Duration) error {
	done := make(chan struct{}, 1)
	cancel := make(chan struct{}, 1)
	go func() {
		d := false
		for {
			if !d {
				select {
				case <-cancel:
					return
				default:
					f()
					d = true
				}
			}
			if d {
				break
			}
		}
		done <- struct{}{}
	}()
	select {
	case <-done:
	case <-time.After(d):
		cancel <- struct{}{}
		return fmt.Errorf("exec func exceedTime error, %v time out", d)
	}
	return nil
}
