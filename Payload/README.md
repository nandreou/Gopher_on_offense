# Gopher_on_offense

This is a Reverse shell Payload writen in GO. The server.go is running on the attacker side and the target.go - target_windows.go are running on target side.
It runs basic commands like ls,pwd,whoami etc. and changes directories, Downloads, and Uploads file to-from the target.

target.go, Is for linux and target_windows.go is for windows.
In windows it may be a little buggy you may need to give a command like "dir" more than one time to get the cmd output.

This is written only for educational and ethical purposes. 
