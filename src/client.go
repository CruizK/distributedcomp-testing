package client

import (
	"fmt"
	"log"
	"net"
)

//const url = "157.230.52.239:4040"
const url = "localhost:4040"

func main() {
	fmt.Println("Connecting to server...")
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}

	fmt.Println(conn.RemoteAddr().String())

	conn.Write([]byte("Hello this is commander hank to ground control\n"))

	for {

	}
}
