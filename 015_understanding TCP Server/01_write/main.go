package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")  // Set up the Listener
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {                                   //Create an eternal loop to accept connections
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "well, I hope!")

		conn.Close()
	}
}
