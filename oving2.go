package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int = 0

func thread1(ch chan int){
	for j := 0; j < 1000000; j++{
		i :=<-ch
		i++
		ch <-i
	}
}

func thread2(ch chan int){
	for j := 0; j < 1000000; j++{
		i := <-ch
		i--
		ch <-i
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 
	ch :=make(chan int, 1)
	ch <- 0

	go thread1(ch)
	//time.Sleep(100*time.Millisecond)

	go thread2(ch)
	time.Sleep(100*time.Millisecond)
  	Println(i)
}