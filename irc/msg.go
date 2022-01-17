// Copyright (c) 2016-2019 Daniel Oaks <daniel@danieloaks.net>
// Copyright (c) 2018-2019 Shivaram Lingamneni <slingamn@cs.stanford.edu>
package irc

import (
	"strings"
)

type Msg struct {
	Prefix string
	Cmd    string
	Params []string
}

func NewMsg(prefix, cmd string, params ...string) Msg {
	return Msg{Prefix: prefix, Cmd: cmd, Params: params}
}

func Parse(line string) (msg Msg, _ error) {
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	line = strings.TrimLeft(line, " ")
	if line[0] == ':' {
		end := strings.IndexByte(line, ' ')
		msg.Prefix = line[1:end]
		line = line[end+1:]
	}
	end := strings.IndexByte(line, ' ')
	if end == -1 {
		end = len(line)
	}
	msg.Cmd = strings.ToUpper(line[:end])
	line = line[end+1:]
	for {
		line = strings.TrimLeft(line, " ")
		if len(line) == 0 {
			break
		}
		if line[0] == ':' {
			msg.Params = append(msg.Params, line[1:])
			break
		}
		space := strings.IndexByte(line, ' ')
		if space == -1 {
			msg.Params = append(msg.Params, line)
			break
		}
		msg.Params = append(msg.Params, line[:space])
		line = line[space+1:]
	}
	return msg, nil
}
