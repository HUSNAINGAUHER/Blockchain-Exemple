package assignment02IBC
import (
"log"
"net"
"fmt"
"strconv"
)


var Port int

func Server(portNumber string,numNodes int) {
	var ln net.Listener
	var err error
	if numNodes>0 {
	ln, err = net.Listen("tcp",":"+portNumber )
	fmt.Println("Server Stated at Port Number",ln.Addr().(*net.TCPAddr).Port)
	} else {
	ln, err = net.Listen("tcp" ,":")
	
	


	fmt.Println("Server Stated at Port Number",ln.Addr().(*net.TCPAddr).Port)
	}

		Port=ln.Addr().(*net.TCPAddr).Port
	
	if numNodes>0 {
		fmt.Println("Server Stated at Port Number",portNumber)
		Conns:=make([]net.Conn,0)
		if err != nil {
		log.Fatal(err)
		}
		i:=0
		for i<numNodes{
		i=i+1
		conn, err := ln.Accept()
		if err != nil {
		log.Println(err)
		continue
		}
		log.Println("Connected to ",i," peers")
		Conns=append(Conns,conn)	
		}

		i=0
		for i<numNodes{
		go handleConnection(Conns[i])
		i=i+1
		}
	} 

	if err != nil {
		log.Fatal(err)
		}

	i:=0
	for {
			i=i+1
			conn, err := ln.Accept()
			if err != nil {
			log.Println(err)
			continue
			}
			log.Println("Connected to ",i," peers")
			go handleConnection(conn)
			}
}
func handleConnection(conn net.Conn) {
	recvdSlice := make([]byte, 11)
	conn.Read(recvdSlice)
	fmt.Println(string(recvdSlice))

	fmt.Println(Port)
	conn.Write([]byte(string(Port)))
}



func Client(portNumber string) {
conn, err := net.Dial("tcp", "localhost:"+portNumber)
if err != nil {

// handle error

}

for {
fmt.Println(Port)
a:=strconv.Itoa(Port)
fmt.Println(a)
conn.Write([]byte(string(a)))

recvdSlice := make([]byte, 11)
conn.Read(recvdSlice)
fmt.Println("dd"+string(recvdSlice))

}
	
}