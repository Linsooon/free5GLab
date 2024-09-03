package lab0

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type listenerInterface func(string, int, handlerInterface)

type handlerInterface func(conn net.Conn)

func TCPListener(host string, port int, handler handlerInterface) {
	listener, err := net.Listen("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine to allow multiple connections
		go TCPHandler(conn)
	}
}

func TCPHandler(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received from %s: %s\n", clientAddr, message)

		// Send a response back to the client
		response := message + "\n"
		conn.Write([]byte(response))
	}
}
