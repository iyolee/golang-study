package web

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// 	"strings"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:3000")

// 	if err != nil {
// 		fmt.Println("Error dialing", err.Error())
// 		return
// 	}

// 	inputReader := bufio.NewReader(os.Stdin)

// 	fmt.Println("What is your name")
// 	clientName, _ := inputReader.ReadString('\n')
// 	trimmedClient := strings.Trim(clientName, "\n")
// 	for {
// 		fmt.Println("What to send to the server? Type Q to quit.")
// 		input, _ := inputReader.ReadString('\n')
// 		trimmedInput := strings.Trim(input, "\r\n")
// 		if trimmedInput == "Q" {
// 			return
// 		}
// 		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
// 	}
// }
