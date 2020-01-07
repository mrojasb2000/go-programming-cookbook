package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn) {
	// set up a reader on conn (an io.Reader)
	reader := bufio.NewReader(conn)

	// grab the first line of data encountered
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", data)
		return
	}
	// print then send back the data
	fmt.Printf("Received: %s", data)

	conn.Write([]byte(strings.ToUpper(data)))

	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("listering on; %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting conection: %s\n", err.Error())
			// if there's an error try again
			continue
		}
		// handle this asynchronously
		// potencially a good use-case
		// for a worker pool
		go echoBackCapitalized(conn)
	}
}
