package main

import (
	"DIV-01/real-time-forum/internal/apiserver"
	"DIV-01/real-time-forum/internal/store/sqlstore"
	"log"
)

func main() {

	sqlO := &sqlstore.Options{
		Address: "./local.db",
	}

	st, err := sqlstore.Start(sqlO)
	defer st.Close()
	if err != nil {
		log.Fatal(err)
	}
	srv := apiserver.NewServer(st)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
