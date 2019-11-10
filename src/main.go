package main

import (
	"fmt"
	"net"
	"web"
)

func main() {
	// echo.go
	// http.HandleFunc("/", web.Echo)
	// err := http.ListenAndServe(":3000", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// tcp.go
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go web.DoServer(conn)
	}
}
