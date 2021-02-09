package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleClientConnection(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, "hello\n")
}

func main() {
	router := httprouter.New()
	router.GET("/client/connect", handleClientConnection)

	l, err := net.Listen("tcp", ":4040")

	if err != nil {
		fmt.Println("Error listening:", err.Error())

	}

	defer l.Close()

	//log.Fatal(http.ListenAndServe(":8080", router))

	for {
		// Holds execution until a connection comes
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Connecting", err.Error())
			return
		}

		fmt.Println("Client connected")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Reading client msg")
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	log.Println("Client Message:", string(buffer[:len(buffer)-1]))

	conn.Write(buffer)

	handleConnection(conn)
}
