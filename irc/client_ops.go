package irc

import "log"

func (cl *Client) SetNick(nick string) error {
	oldNick := cl.nick

	if _, ok := cl.srv.Nicks[nick]; ok {
		return ERR_NICKNAMEINUSE
	}

	cl.srv.Mu.Lock()
	delete(cl.srv.Nicks, cl.nick)
	cl.srv.Nicks[nick] = struct{}{}
	cl.nick = nick
	cl.srv.Mu.Unlock()

	if cl.regged {
		err := cl.sendLine(oldNick, "NICK", cl.nick)
		if err != nil {
			return err
		}
		log.Print("SEND NICK TO ALL CHANNELS HERE")
		log.Print(cl.user, " set nick from ", oldNick, " to ", cl.nick)
	} else {
		log.Print(cl.Addr(), " set nick to ", cl.nick)
		if cl.user != "" && cl.realname != "" {
			cl.regged = true
			cl.Welcome()
		}
	}

	return nil
}

func (cl *Client) Register(user string, realname string) {
	if cl.regged {
		panic("send ERR_ALREADYREGISTERED")
	}
	cl.user = user
	cl.realname = realname
	if cl.nick != "" {
		cl.regged = true
		cl.Welcome()
	}
}
