package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
)

var channel = "#1125"

func main() {
	con := irc.IRC("murgrabia", "murgrabia")
	err := con.Connect("krakow.irc.pl:6667")

	if err != nil {
		fmt.Println("Connection failed")
		return
	}

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(channel)
	})

	con.AddCallback("JOIN", func(e *irc.Event) {
		con.Privmsg(channel, "Sup!")
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		if strings.Contains(e.Message(), "!quit") {
			con.SendRaw("QUIT :I quit")
			con.Quit()
		}
	})

	con.Loop()
}
