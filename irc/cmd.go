package irc

import "fmt"

const (
	passMsgTempl   = "PASS %s"
	nickMsgTempl   = "NICK %s"
	userMsgTempl   = "USER %s * * :%s"
	joinMsgTempl   = "JOIN %s"
	partMsgTempl   = "PART %s"
	quitMsgTempl   = "QUIT %s"
	privMsgTempl   = "PRIVMSG %s :%s"
	noticeMsgTempl = "NOTICE %s :%s"
)

func passMsg(password string) string {
	return fmt.Sprintf(passMsgTempl, password)
}

func nickMsg(username string) string {
	return fmt.Sprintf(nickMsgTempl, username)
}

func userMsg(username, realname string) string {
	return fmt.Sprintf(userMsgTempl, username, realname)
}

func joinMsg(channel, password string) string {
	cmd := channel

	if password != "" {
		cmd += " " + password
	}
	return fmt.Sprintf(joinMsgTempl, cmd)
}

func partMsg(channel string) string {
	return fmt.Sprintf(partMsgTempl, channel)
}

func quitMsg(msg string) string {
	return fmt.Sprintf(quitMsgTempl, msg)
}

func privMsg(target, msg string) string {
	return fmt.Sprintf(privMsgTempl, target, msg)
}

func noticeMsg(target, msg string) string {
	return fmt.Sprintf(noticeMsgTempl, target, msg)
}
