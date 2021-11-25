package irc

import (
	"net"
)

// Conn wraps a network connection in case we need to add more context to it
type Conn struct {
	net.Conn
}
