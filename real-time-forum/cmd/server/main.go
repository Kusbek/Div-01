package main

import (
	"DIV-01/real-time-forum/internal/apiserver"
	"log"
)

func main() {
	o := apiserver.Options{
		Address: ":8082",
	}

	if err := apiserver.Start(o); err != nil {
		log.Fatal(err)
	}
}
