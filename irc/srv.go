package irc

import (
	"log"
	"net"
)

type Srv struct {
	ls net.Listener
}

func NewSrv() (*Srv, error) {
	ls, err := net.Listen("tcp", ":6667")
	if err != nil {
		return nil, err
	}
	return &Srv{ls}, nil
}

func (srv *Srv) Addr() net.Addr { return srv.ls.Addr() }

// Listen does not return
func (srv *Srv) Listen() {
	ch := make(chan error)
	go func() {
		for {
			s, err := srv.ls.Accept()
			if err != nil {
				ch <- err
				continue
			}
			c := &Conn{s}
			log.Printf("accepted connection from %s", c.RemoteAddr().String())
			go srv.newClient(c)
		}
	}()
	for {
		log.Print((<-ch).Error())
	}
}

func (srv *Srv) newClient(c *Conn) {
	c.Write([]byte(""))
	c.Close()
}
