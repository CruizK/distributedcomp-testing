package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Connecting to server...")
	conn, err := net.Dial("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}

	fmt.Println(conn.RemoteAddr().String())
}
