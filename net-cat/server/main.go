package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type client struct {
	connection *net.Conn
	name       string
}

type chatServer struct {
	clients []*client
	mutex   sync.Mutex
}

func main() {
	startServer()
}
func startServer() {
	addr := "127.0.0.1:8989"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening for connections on %s", listener.Addr().String())
	chat := &chatServer{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
		} else {
			if len(chat.clients) < 2 {
				_ = threeWayHandshake(&conn, chat)
			} else {
				fmt.Fprintf(conn, "Sorry, bro, but chat is polnyi$")
			}
			fmt.Println(chat.clients)
		}
	}
}

func threeWayHandshake(conn *net.Conn, chat *chatServer) *client {
	fmt.Fprintf(*conn, "Atyndy jaz, sumelek: $")
	reader := bufio.NewReader(*conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	client := &client{connection: conn, name: name}
	chat.clients = append(chat.clients, client)
	fmt.Fprintf(os.Stdout, name+" has joined our chat...\n")
	return client
}
