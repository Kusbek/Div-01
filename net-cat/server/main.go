package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type client struct {
	connection *net.Conn
	name       string
}

type message struct {
	client    *client
	message   string
	timestamp string
}

type chatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
	distr    chan *message
}

const (
	LIMIT int = 2
)

func main() {
	args := os.Args[1:]
	var port int
	var err error
	if len(args) == 1 {
		port, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Print("Invalid input:")
			fmt.Println(err)
			return
		}
	} else if len(args) < 1 {
		port = 8989
	} else {
		fmt.Print("Too many parameters")
		return
	}

	chat := &chatServer{mutex: &sync.Mutex{}, distr: make(chan *message, LIMIT)}
	chat.startServer(port)
}
func (self *chatServer) startServer(port int) {
	addr := fmt.Sprintf("localhost:%d", port)
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
			if len(self.clients) < LIMIT {
				client := self.threeWayHandshake(&conn)
				go client.serve(self)
				go self.broadcast()
			} else {
				fmt.Fprintf(conn, "Sorry, bro, but chat is polnyi")
			}
		}
	}
}
func (self *client) serve(chat *chatServer) {
	reader := bufio.NewReader(*self.connection)
	for {
		fmt.Fprintf(*self.connection, fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), self.name))
		msg, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Printf("detected closed client: %s\n", self.name)
			self.remove(chat)
			chat.goodbye(self)
			break
		}
		time := time.Now().Format("2006-01-02 15:04:05")
		chat.distr <- &message{client: self, message: msg, timestamp: time}
		fmt.Printf("[%s][%s]: %s", time, self.name, msg)
	}
}

func (self *client) remove(chat *chatServer) {
	chat.mutex.Lock()
	defer chat.mutex.Unlock()
	for i, v := range chat.clients {
		if v == self {
			chat.clients = append(chat.clients[:i], chat.clients[i+1:]...)
		}
	}
}

func (self *chatServer) broadcast() {
	for {
		cl := <-self.distr
		self.mutex.Lock()
		for _, client := range self.clients {
			if client != cl.client {
				fmt.Fprintf(*client.connection, fmt.Sprintf("\n[%s][%s]: %s", cl.timestamp, cl.client.name, cl.message))
				fmt.Fprintf(*client.connection, fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.name))
			}
		}
		self.mutex.Unlock()
	}
}

func (self *chatServer) introduce(cl *client) {
	for _, client := range self.clients {
		fmt.Fprintf(*client.connection, "\n"+cl.name+" has joined our chat...\n")
		fmt.Fprintf(*client.connection, fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.name))
	}
}
func (self *chatServer) goodbye(cl *client) {
	for _, client := range self.clients {
		fmt.Fprintf(*client.connection, "\n"+cl.name+" has left our chat...\n")
		fmt.Fprintf(*client.connection, fmt.Sprintf("[%s][%s]:", time.Now().Format("2006-01-02 15:04:05"), client.name))
	}
}

func (self *chatServer) threeWayHandshake(conn *net.Conn) *client {
	file, _ := os.Open("linux.txt")
	text, _ := ioutil.ReadAll(file)
	fmt.Fprintf(*conn, string(text))
	fmt.Fprintf(*conn, "Atyndy jaz, sumelek: ")
	reader := bufio.NewReader(*conn)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	client := &client{connection: conn, name: name}
	self.mutex.Lock()
	self.introduce(client)
	self.clients = append(self.clients, client)
	defer self.mutex.Unlock()
	fmt.Fprintf(os.Stdout, name+" has joined our chat...\n")
	return client
}
