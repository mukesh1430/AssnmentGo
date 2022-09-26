package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

var port = "0.0.0.0:9001"

var counter int = 0

func echoMessage(conn net.Conn) {
	r := bufio.NewReader(conn)
	for {
		message, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("Error ", err)

		}
		conn.Write([]byte(fmt.Sprintf("%s   ", time.Now())))
		conn.Write([]byte("From Server :"))

		conn.Write(message)

		fmt.Println("Message Receied From Client : ", string(message), conn.RemoteAddr())
		fmt.Printf("No of Client Conected %d \n", counter)
	}

}
func main() {
	fmt.Println("Start the server ")
	ln, err := net.Listen("tcp", port)
	for {
		conn, _ := ln.Accept()
		fmt.Println(" Connected Client : ", conn.RemoteAddr())
		counter++
		fmt.Printf("No of Client Conected %d \n", counter)
		if err != nil {
			fmt.Println("Error ", err)
			continue
		}

		go echoMessage(conn)

	}

}
