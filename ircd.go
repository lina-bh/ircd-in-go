package main

import (
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/linabeee/ircd-in-go/irc"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("what the fuck:", err.Error())
	}
	conf := irc.SrvConf{
		Hostname: hostname,
	}
	srv, err := irc.NewSrv(&conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Print(srv.Addr().String())
	errCh := srv.Listen(true)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)
	for {
		select {
		case <-sigCh:
			pprof.Lookup("goroutine").WriteTo(log.Default().Writer(), 1)
			os.Exit(127)
		case err := <-errCh:
			log.Println(err.Error())
		}
	}
}
