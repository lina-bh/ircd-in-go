package irc

import (
	"fmt"
)

func (cl *Client) Welcome() (err error) {
	err = cl.sendLine(cl.srv.Hostname, "001", cl.nick, fmt.Sprint("Welcome to IRC, ", cl.nick, "!"))
	if err != nil {
		return err
	}
	err = cl.sendLine(cl.srv.Hostname, "002", cl.nick, "You're flying linabee/ircd-in-go@0")
	if err != nil {
		return err
	}
	time := cl.srv.Began.String()
	err = cl.sendLine(cl.srv.Hostname, "003", cl.nick, fmt.Sprint("This server created ", time))
	if err != nil {
		return err
	}
	err = cl.myInfo()
	if err != nil {
		return err
	}
	err = cl.iSupport()
	if err != nil {
		return err
	}
	return
}

func (cl *Client) myInfo() error {
	// get these from srv
	version := "0"
	umodes := "i"
	cmodes := "n"
	// amodes := "b"
	return cl.sendLine(cl.srv.Hostname, "004", cl.nick, cl.srv.Hostname, version, umodes, cmodes)
}

func (cl *Client) iSupport() error {
	// be reasonable about this
	return cl.sendLine(cl.srv.Hostname, "005", cl.nick, "CASEMAPPING=ascii", "are supported by this server")
}
