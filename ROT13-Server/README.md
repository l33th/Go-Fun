# This is a rot13 algorithm server

# Running the server:
	## Linux:
		- download this repository or copy the code
		- when inside the directory where the rot13 algorithm is run the folling command in a terminal:
			go run rot13Server.go
		
		- open a new terminal and type the following:
			* telnet localhost 8080
			* type 'hello', the output should be 'uryyb'
			* if you enter 'uryyb', the output should be 'hello'
