package main

import (
	"fmt"
	"net"
	"time"
)

func client(){
	conn, err := net.Dial("tcp", "localhost:20003")
	if err != nil {
		// handle error
	}
	conn.Write([]byte("Hey amigo!"))

	buf := make([]byte, 1024)

	ln, err := conn.Read(buf)
	if err!=nil {
		//
	}

	fmt.Println(string(buf[0:ln]))




}

func server(){
	ln, err := net.Listen("tcp", "localhost:20003")
	if err != nil {
		fmt.Println(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}

		go handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024)

	length, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(buf[0:length]))

	conn.Write([]byte("Message received"))

	conn.Close()

}


func main(){
	go server()
	go client()
	time.Sleep(time.Second*15)
}