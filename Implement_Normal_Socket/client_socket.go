package main

import (
	"net"
	"../Implement_Normal_Socket/socket"
	"fmt"
)


func HandleWrite(conn net.Conn, msg chan string) {
	_, e := conn.Write([]byte("hello"))
	socket.Checkerror(e)
	msg <- "Sent"
}

func HandleRead(conn net.Conn, msg chan string) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	socket.Checkerror(err)
	fmt.Println(string(buffer[:n]))
	msg <- "Recieved"
}

func main()  {
	go socket.Run()
	conn, err := net.Dial("tcp", "localhost" + ":" + "2333")
	socket.Checkerror(err)

	defer conn.Close()

	fmt.Println("Connecting to " + "localhost" + ":" + "2333")
	msg := make(chan string)
	go HandleWrite(conn, msg)
	go HandleRead(conn, msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
