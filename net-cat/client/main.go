package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	Start()
}
func Start() {
	c := Dial()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go listenClient(&c, &wg)
	wg.Add(1)
	go listenServer(&c, &wg)
	wg.Wait()
}

func Dial() net.Conn {
	CONNECT := "127.0.0.1:8989"
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	message, _ := bufio.NewReader(c).ReadString('$')
	message = strings.TrimSuffix(message, "$")
	fmt.Fprintf(os.Stdout, message+" ")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Fprintf(c, name)
	return c
}

func listenServer(c *net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	serverReader := bufio.NewReader(*c)
	for {
		message, err := serverReader.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Printf("detected closed LAN connection")
			break
		}
		fmt.Print(message)
	}
}

func listenClient(c *net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(*c, text)
	}
}
