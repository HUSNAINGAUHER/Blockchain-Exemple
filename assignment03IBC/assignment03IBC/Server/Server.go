package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	a1 "github.com/HUSNAINGAUHER/assignment01IBC"
)

var chainHead *a1.Block
var wg sync.WaitGroup
var command chan string
var To string
var Amount string
var Miner int

//Server of Peer to peer Network
// portNumber on which this serevr listen
// numNodes in case of satoshi num of nodes to be connected

type Node struct {
	Peers []string
}

var Peer Node

func Start_Server(portNumber string, numNodes int) int {

	if numNodes > 0 {
		Peer.Peers = append(Peer.Peers, string(portNumber))
		ln, err := net.Listen("tcp", ":"+portNumber)
		if err != nil {
			log.Println(err)
		}
		go Server(portNumber, numNodes, ln)
		return ln.Addr().(*net.TCPAddr).Port
	} else {
		ln, err := net.Listen("tcp", ":")
		go Server(portNumber, numNodes, ln)
		if err != nil {
			log.Println(err)
		}
		return ln.Addr().(*net.TCPAddr).Port
	}
}

func Server(portNumber string, numNodes int, ln net.Listener) {

	chainHead = a1.InsertBlock("TO SANTOSHI 100 COINS Gensis Block", nil)
	if numNodes > 0 {
		Satoshi(ln, numNodes)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)

	}

}

func Satoshi(ln net.Listener, num int) {
	i := 0
	Conns := make([]net.Conn, 0)
	for i < num {

		conn, err := ln.Accept()
		chainHead = a1.InsertBlock("TO SANTOSHI 100 COINS", chainHead)
		if err != nil {
			log.Println(err)
		}

		Conns = append(Conns, conn)
		i = i + 1
	}

	i = 0
	for i < num {
		go handleConnection(Conns[i])
		i = i + 1
	}

}

func handleConnection(conn net.Conn) {

	recvdSlice := make([]byte, 11)
	conn.Read(recvdSlice)
	Port := ""

	// ... Convert back into a string from rune slice.
	recvdSlice = recvdSlice[0:6]

	if string(recvdSlice) != "second" {
		Peer.Peers = append(Peer.Peers, string(recvdSlice))
		Port = string(recvdSlice)
		conn.Write([]byte(string("ready")))

		gobEncoder := gob.NewEncoder(conn)
		err := gobEncoder.Encode(chainHead)
		if err != nil {

			log.Println(err)

		}

		fmt.Println(Peer.Peers[0])
		gobEncoder = gob.NewEncoder(conn)
		err = gobEncoder.Encode(Peer.Peers)
		if err != nil {

			log.Println(err)

		}

	} else {

		fmt.Println("Connected")
		go routine(command, &wg, Port)
	}

}

func routine(command <-chan string, wg *sync.WaitGroup, Port string) {
	defer wg.Done()
	var status = "Pause"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work(Port)
			}
		}
	}
}

func work(Port string) {

	if AS(Peer.Peers[Miner]) == AS(Port) {
		fmt.Println("Miner Found")
	}
}

func Client(portNumber string, ServerPort string) {
	conn, err := net.Dial("tcp", "localhost:"+portNumber)
	if err != nil {

		// handle error

	}

	a := ServerPort
	conn.Write([]byte(string(a)))

	recvdSlice := make([]byte, 11)
	conn.Read(recvdSlice)

	var recvdBlock a1.Block
	dec := gob.NewDecoder(conn)
	err = dec.Decode(&recvdBlock)
	if err != nil {

		fmt.Println("jnfjfn")

	}

	chainHead = &recvdBlock
	a1.Display(chainHead)

	var recv []string
	dec = gob.NewDecoder(conn)
	err = dec.Decode(&recv)
	if err != nil {
		fmt.Println("jnfjfn")
	}
	Peer.Peers = recv

	i := 0
	for i < len(Peer.Peers) {

		if AS(ServerPort) == AS(Peer.Peers[i]) {
			fmt.Println("self")
		} else {
			fmt.Println("Connecting to " + Peer.Peers[i])
			go ClientDail(Peer.Peers[i])
		}
		i = i + 1
	}

}

func checkbyte(p []rune) int {

	i := 0
	for i < len(p) {
		if p[i] == 0 {
			return i
		}
		i = i + 1
	}
	return len(p)
}

func AS(ServerPort string) string {
	runes := []rune(ServerPort)
	// ... Convert back into a string from rune slice.
	p := checkbyte(runes)
	fmt.Println(p)
	ServerPort = string(runes[0:p])

	return ServerPort
}

func ClientDail(ServerPort string) {

	runes := []rune(ServerPort)
	// ... Convert back into a string from rune slice.
	p := checkbyte(runes)
	fmt.Println(p)
	ServerPort = string(runes[0:p])
	conn, err := net.Dial("tcp", "localhost:"+ServerPort)

	if err != nil {

		fmt.Println("error")

	}

	conn.Write([]byte(string("second")))

}

func main() {

	command = make(chan string)
	wg.Add(1)

	if len(os.Args) > 2 {
		i1, err := strconv.Atoi(os.Args[2])
		if err != nil {

		}
		Start_Server(os.Args[1], i1)
	}

	if len(os.Args) == 2 {

		Port := Start_Server(os.Args[1], -1)
		Port1 := toString(Port)
		go Client(os.Args[1], Port1)

		//normal case
	}

	command <- "Pause"
	for {

		time.Sleep(1 * time.Second)
		command <- "Pause"

		fmt.Println("Enter any key to make a transaction")
		fmt.Scanln(&To)
		fmt.Println("To whoom")
		fmt.Scanln(&To)
		fmt.Println("Amount")
		fmt.Scanln(&Amount)
		fmt.Println("Select Miner")
		i := 0
		for i < len(Peer.Peers) {
			fmt.Println("Press" + string(i+48) + "for " + Peer.Peers[i])
			i = i + 1
		}

		fmt.Scanln(&Miner)

		time.Sleep(1 * time.Second)
		command <- "Play"
		wg.Wait()

	}

}

func toString(num int) string {

	temp := ""
	for num > 0 {
		temp = string(num%10+48) + temp
		num = num / 10
	}

	return temp
}
