package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Connecting to server...")
	conn, err := net.Dial("tcp", "157.230.52.239:4040")
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}

	fmt.Println(conn.RemoteAddr().String())
}
