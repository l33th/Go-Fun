# This is a rot13 algorithm server

# Running the server:
	## Linux:
		- Open a new terminal
		- git clone https://github.com/l33th/GoFun/ROT13-Server
		- cd ROT13-Server
		- go run rot13Server.go
		- open a new terminal and type the following:
			* telnet localhost 8080
			* type 'hello', the output should be 'uryyb'
			* if you enter 'uryyb', the output should be 'hello'
