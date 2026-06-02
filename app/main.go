package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO: Uncomment the code below to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	req := make([]byte, 4096)
	_, err = conn.Read(req)
	if err != nil {
		log.Printf("error reading request: %s", err)
	}

	parts := strings.Split(string(req), "\r\n")
	// len := len(parts)
	requestSplit := strings.Split(parts[0], " ")
	// method := requestSplit[0]
	path := requestSplit[1]
	// headers := parts[1 : len-2]
	// requestBody := parts[len-1]
	// fmt.Printf("method: %s headers: %s body: %s", method, headers, requestBody)
	var response string
	pathSplit := strings.Split(path, "/")
	fmt.Println(len(pathSplit))
	// if  {
	// 	response = "HTTP/1.1 404 Not Found\r\n\r\n"
	// }
	if len(pathSplit) < 3 {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}
	responseEcho := pathSplit[2]
	// response = "HTTP/1.1 200 OK\r\n\r\n"
	response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(pathSplit), responseEcho)
	fmt.Println(pathSplit[2])
	conn.Write([]byte(response))
}
