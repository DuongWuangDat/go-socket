package main

import (
	"bufio"
	"fmt"
	"strconv"

	"log"
	"net"
	"os"
)

func main() {
	var port string
	for {
		fmt.Println("Room: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		x, erro := strconv.Atoi(scanner.Text())
		if erro == nil {
			port = ":300" + strconv.Itoa(x)
			break
		}

	}
	connection, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	nameReader := bufio.NewScanner(os.Stdin)
	nameReader.Scan()
	nameInput := nameReader.Text()
	//fmt.Println(nameInput)
	go onMessage(connection)
	for {
		msgReader := bufio.NewScanner(os.Stdin)
		msgReader.Scan()
		msg := msgReader.Text()
		if msg == "" {
			break
		}
		//fmt.Println(nameInput)
		msg = fmt.Sprintf("%s: %s\n", (nameInput), (msg))

		_, err = connection.Write([]byte(msg))
		if err != nil {
			break
		}

	}
	connection.Close()
}
func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Print(msg)
	}
}
