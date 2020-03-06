package main

import (
	"bufio"
	"fmt"
	"io"
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

type message struct {
	client  *client
	message string
}

type chatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
	distr    chan *message
}

func main() {
	chat := &chatServer{mutex: &sync.Mutex{}, distr: make(chan *message, 1)}
	chat.startServer()
}
func (self *chatServer) startServer() {
	addr := "127.0.0.1:8989"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening for connections on %s", listener.Addr().String())
	self.listener = listener

	for {
		conn, err := self.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection from client: %s", err)
		} else {
			if len(self.clients) < 2 {
				client := self.threeWayHandshake(&conn)
				go client.serve(self)
				go self.broadcast()
			} else {
				fmt.Fprintf(conn, "Sorry, bro, but chat is polnyi$")
			}
		}
	}
}

func (self *chatServer) broadcast() {
	for {
		cl := <-self.distr
		self.mutex.Lock()
		for _, client := range self.clients {
			if client != cl.client {
				fmt.Fprintf(*client.connection, fmt.Sprintf("%s: %s", cl.client.name, cl.message))
			}
		}
		self.mutex.Unlock()
	}
}

func (self *chatServer) introduce(cl *client) {
	for _, client := range self.clients {
		fmt.Fprintf(*client.connection, cl.name+" has joined our chat...\n")
	}
}

func (self *chatServer) threeWayHandshake(conn *net.Conn) *client {
	fmt.Fprintf(*conn, "Atyndy jaz, sumelek: $")
	reader := bufio.NewReader(*conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	client := &client{connection: conn, name: name}
	self.mutex.Lock()
	self.introduce(client)
	fmt.Println("Finished")
	self.clients = append(self.clients, client)
	defer self.mutex.Unlock()
	fmt.Fprintf(os.Stdout, name+" has joined our chat...\n")
	return client
}

func (self *client) serve(chat *chatServer) {
	reader := bufio.NewReader(*self.connection)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Printf("detected closed client: %s", self.name)
			break
		}
		chat.mutex.Lock()
		chat.distr <- &message{client: self, message: msg}
		chat.mutex.Unlock()
		fmt.Printf("%s: %s", self.name, msg)
	}
}
