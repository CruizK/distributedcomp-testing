package main

import (
	"fmt"
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

	l, err := net.Listen("tcp", "localhost:4040")

	if err != nil {
		fmt.Println("Error listening:", err.Error())

	}

	defer l.Close()

	//log.Fatal(http.ListenAndServe(":8080", router))

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Connecting", err.Error())
			return
		}

		fmt.Println("Client connected")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected")
	}
}
