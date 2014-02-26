package main

// 伪代码
import (
	"fmt"
)

func main() {
	pool := NewFuncPool(5)
	for i := 0; i < 1000; i++ {
		func(n int) {
			pool.Run(func() {
				fmt.Println(n)
			})
		}(i)
	}
	pool.Wait()
	fmt.Println("OK")
}

type FuncPool struct {
	buffer chan int
}

func NewFuncPool(size int) (f *FuncPool) {
	f = &FuncPool{
		buffer: make(chan int, size),
	}
	for i := 0; i < cap(f.buffer); i++ {
		f.buffer <- i
	}
	return
}

func (f *FuncPool) Run(fn func()) {
	i := <-f.buffer
	go func() {
		fn()
		f.buffer <- i
	}()
}

func (f *FuncPool) Wait() {
	for i := 0; i < cap(f.buffer); i++ {
		<-f.buffer
	}
}
