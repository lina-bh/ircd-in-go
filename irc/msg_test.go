package irc_test

import (
	"testing"

	. "github.com/linabeee/ircd-in-go/irc"
)

func TestCapLs(t *testing.T) {
	line := "CAP LS\r\n"
	msg, err := Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Cmd != "CAP" {
		t.Fatal("msg.cmd != \"CAP\"")
	}
	if len(msg.Params) != 1 || msg.Params[0] != "LS" {
		t.Fatal("msg.params[0] != \"LS\"")
	}
}

func TestNick(t *testing.T) {
	line := "NICK nick"
	msg, err := Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Cmd != "NICK" {
		t.Fatal("msg.cmd != \"NICK\"")
	}
	if len(msg.Params) != 1 || msg.Params[0] != "nick" {
		t.Fatal("msg.params[0] != \"nick\"")
	}
}

func TestUser(t *testing.T) {
	line := "USER username username hostname :Realname"
	msg, err := Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Cmd != "USER" {
		t.Fatal("msg.cmd != \"USER\"")
	}
	if len(msg.Params) != 4 {
		t.Fatal("len(msg.params) != 4")
	}
	if msg.Params[0] != "username" {
		t.Fatal("msg.params[0] != \"username\"")
	}
	if msg.Params[3] != "Realname" {
		t.Fatal("msg.params[3] != \"Realname\"")
	}
}

func TestJoin(t *testing.T) {
	line := ":nick!user@127.0.0.1 JOIN #chan\r\n"
	msg, err := Parse(line)
	if err != nil {
		t.Fatal(err)
	}
	if msg.Prefix != "nick!user@127.0.0.1" {
		t.Fatal("msg.prefix != \"nick!user@127.0.0.1\"")
	}
	if msg.Cmd != "JOIN" {
		t.Fatal("msg.cmd != \"JOIN\"")
	}
	if len(msg.Params) != 1 {
		t.Fatal("len(msg.params) != 1")
	}
	if msg.Params[0] != "#chan" {
		t.Fatal("msg.params[0] != \"#chan\"")
	}
}
