package main

import (
	"log"

	"github.com/linabeee/ircd-in-go/irc"
)

func main() {
	srv, err := irc.NewSrv()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(srv.Addr().String())
	srv.Listen()
}
