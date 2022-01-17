package irc

type IRCError struct {
	num  int
	name string
	msg  string
}

var (
	ERR_NONICKNAMEGIVEN = &IRCError{431, "ERR_NONICKNAMEGIVEN", "no nickname"}
	ERR_NICKNAMEINUSE   = &IRCError{433, "ERR_NICKNAMEINUSE", "nickname already used"}
)

func (err *IRCError) Error() string {
	return err.name
}
