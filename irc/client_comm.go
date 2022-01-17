package irc

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

var (
	ErrNotUtf8 error = errors.New("incoming data from client isn't UTF-8")
)

func (cl *Client) Addr() string { return cl.conn.RemoteAddr().String() }

func (cl *Client) readMsg() (msg Msg, _ error) {
	buf, err := cl.reader.ReadBytes('\n')
	if err != nil {
		return msg, err
	}
	if !utf8.Valid(buf) {
		return msg, ErrNotUtf8
	}
	line := string(buf)
	msg, err = Parse(line)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (cl *Client) sendLine(prefix, cmd string, params ...string) error {
	var line string

	if prefix != "" {
		line += ":" + prefix + " "
	}
	line += cmd + " "
	var trailer string

	if len(params) != 0 && strings.IndexByte(params[len(params)-1], ' ') != -1 {
		trailer = params[len(params)-1]
		if len(params) > 1 {
			params = params[:len(params)-1]
		} else {
			params = []string{}
		}
	}

	for _, param := range params {
		line += param + " "
	}

	if trailer != "" {
		line += ":" + trailer
	}

	line = strings.TrimSuffix(line, " ")
	line += "\r\n"

	_, err := cl.conn.Write([]byte(line))
	if err != nil {
		return err
	}
	log.Printf("-> %s '%s'\n", cl.Addr(), line[:len(line)-2])
	return nil
}

func (cl *Client) SendError(err *IRCError) {
	simple := func(errorMessage string) {
		cl.sendLine(cl.srv.Hostname, fmt.Sprint(err.num), cl.nick, errorMessage)
	}
	switch err {
	case ERR_NONICKNAMEGIVEN:
		simple("no nickname")
	case ERR_NICKNAMEINUSE:
		simple("nickname already used")
	}
}
