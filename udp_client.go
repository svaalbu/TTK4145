package main

import (
	"fmt"
	"net"
	"time"
)



func receiver2() {
    ServerAddr,err := net.ResolveUDPAddr("udp4",":20003")
    if err != nil {
		fmt.Println(err)
	}

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp4", ServerAddr)
    if err != nil {
		fmt.Println(err)
	}

    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}



func broadcast() {
	time.Sleep(time.Second*1)

	ServerAddr, err := net.ResolveUDPAddr("udp4","255.255.255.255:20003")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp4", nil, ServerAddr)
	if err != nil {
		fmt.Println(err)
	}


	conn.Write([]byte("dette er bredkast"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {


	go broadcast()
	go receiver2()
	time.Sleep(time.Second*5)

}











/*func receiver() {

	conn, err := net.Listen("udp", ":20003")
	if err != nil {
		fmt.Println(err)
	}

	for i<10 {
		conn.
	}

}


func sender() int{

	ServerAddr,err := net.ResolveUDPAddr("udp",":20003")
    if err != nil {
		fmt.Println(err)
	}

	//var IPAddr string = ""

	//IPAddr = string(ServerAddr.IP[:])

	fmt.Printf(string(ServerAddr.IP))	

	conn, err := net.Dial("udp", ":20003")

	if err != nil {
		return 0
	}

	defer conn.Close()
	var i int = 0

	for i<10 {

		i++
		conn.Write([]byte("Hello from sender"))

		time.Sleep(time.Second*1)
	}

	return 1
}
*/