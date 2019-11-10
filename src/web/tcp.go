package web

import (
	"fmt"
	"net"
)

// DoServer tcp
func DoServer(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
