package socket

import (
	"flag"
	"net"
	"io"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "2333", "port")

func Checkerror(err error) {
	if err != nil{
		panic(err)
	}
}

func handleconn(conn net.Conn) {
	defer conn.Close()

	io.Copy(conn, conn)
}

func Run()  {
	l, err := net.Listen("tcp", *host + ":" + *port)
	Checkerror(err)

	defer l.Close()

	for {
		conn, err := l.Accept()
		Checkerror(err)

		go handleconn(conn)
	}
}
