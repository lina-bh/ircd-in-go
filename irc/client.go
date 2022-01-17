package irc

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Client struct {
	srv      *Srv
	conn     net.Conn
	reader   *bufio.Reader
	nick     string
	user     string
	realname string
	regged   bool
}

func NewClient(srv *Srv, s net.Conn) *Client {
	reader := bufio.NewReaderSize(s, 512)
	return &Client{srv: srv, conn: s, reader: reader}
}

func (cl *Client) Run() {
	for {
		msg, err := cl.readMsg()
		if err == io.EOF {
			cl.Disconnect()
			log.Print(cl.Addr(), " dropped")
			return
		} else if err != nil {
			log.Print(err.Error())
			continue
		}
		log.Printf("<- %s %+v\n", cl.Addr(), msg)
		if cmd, ok := cmds[msg.Cmd]; ok {
			err := cmd(cl, msg.Prefix, msg.Params)
			if err != nil {
				if err == io.EOF {
					cl.Disconnect()
				}
				if ircerr, ok := err.(*IRCError); ok {
					cl.SendError(ircerr)
				}
			}
		}
	}
}

func (cl *Client) Disconnect() {
	if err := cl.conn.Close(); err != nil {
		log.Print(err.Error())
	}
	delete(cl.srv.Nicks, cl.nick)
}
