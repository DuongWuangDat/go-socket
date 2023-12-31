package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	//roomID  []int
	conns   []net.Conn
	connCh  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh   = make(chan string)
)

func main() {
	var port string
	for {
		fmt.Println("Add new room: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		x, erro := strconv.Atoi(scanner.Text())
		if erro == nil {
			//	roomID = append(roomID, x)
			port = ":300" + strconv.Itoa(x)
			break
		}

	}

	sever, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			conn, err := sever.Accept()
			if err != nil {
				log.Fatal(err)
			}
			conns = append(conns, conn)
			connCh <- conn
		}
	}()

	for {
		select {
		case conn := <-connCh:
			go onMessage(conn)
		case msg := <-msgCh:
			fmt.Print(msg)
		case conn := <-closeCh:
			fmt.Println("Exit")
			removeConn(conn)
		}
	}

}
func removeConn(conn net.Conn) {
	var i int
	for i := range conns {
		if conns[i] == conn {
			break
		}
	}

	conns = append(conns[i:], conns[:i+1]...)
}
func publishMsg(conn net.Conn, msg string) {
	for i := range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}
}
func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msgCh <- msg
		publishMsg(conn, msg)
	}
	closeCh <- conn
}
