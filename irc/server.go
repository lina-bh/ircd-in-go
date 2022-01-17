package irc

import (
	"log"
	"net"
	"sync"
	"time"
)

type Srv struct {
	Mu       sync.Mutex
	ls       net.Listener
	Nicks    map[string]struct{}
	Hostname string
	Began    time.Time
}

type SrvConf struct {
	Hostname string
}

func NewSrv(conf *SrvConf) (*Srv, error) {
	if conf == nil {
		conf = &SrvConf{Hostname: "server"}
	}
	ls, err := net.Listen("tcp", ":6667")
	if err != nil {
		return nil, err
	}
	return &Srv{
		ls:       ls,
		Hostname: conf.Hostname,
		Nicks:    map[string]struct{}{},
		Began:    time.Now(),
	}, nil
}

func (srv *Srv) Addr() net.Addr { return srv.ls.Addr() }

func (srv *Srv) Listen(fork bool) chan error {
	ch := make(chan error)
	go func() {
		for {
			s, err := srv.ls.Accept()
			if err != nil {
				ch <- err
				continue
			}
			log.Print(s.RemoteAddr().String(), " connected")
			cl := NewClient(srv, s)
			go cl.Run()
		}
	}()
	if fork {
		return ch
	}
	for {
		log.Print((<-ch).Error())
	}
}
