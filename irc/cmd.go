package irc

type Cmd func(cl *Client, prefix string, params []string) error

var (
	cmds map[string]Cmd = map[string]Cmd{
		"NICK": Nick,
		"USER": User,
		"PING": Ping,
	}
)

func Nick(cl *Client, _ string, params []string) error {
	var nick string
	if len(params) != 1 {
		return ERR_NONICKNAMEGIVEN
	}
	nick = params[0]
	err := cl.SetNick(nick)
	return err
}

func User(cl *Client, _ string, params []string) error {
	var (
		user     string
		realname string
	)
	if len(params) < 4 {
		for i := len(params); i < 4; i++ {
			params = append(params, "")
		}
	}
	user = params[0]
	if user == "" {
		panic("need to send ERR_NEEDMOREPARAMS")
	}
	realname = params[3]
	if realname == "" {
		realname = "loser"
	}
	cl.Register(user, realname)
	return nil
}

func Ping(cl *Client, prefix string, params []string) error {
	token := params[0]
	return cl.sendLine("", "PONG", cl.srv.Hostname, token)
}
