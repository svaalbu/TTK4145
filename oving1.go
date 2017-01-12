package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i int = 0

func thread1(ch chan int){
	for j := 0; j < 1000000; j++{
		i++
	}
}

func thread2(ch chan int){
	for j := 0; j < 1000000; j++{
		i--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 

	go thread1()
	//time.Sleep(100*time.Millisecond)

	go thread2()
	time.Sleep(100*time.Millisecond)
  	Println(i)
}
