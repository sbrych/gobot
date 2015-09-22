package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
	"io"
	"io/ioutil"
	"bufio"
	"os"
	"encoding/json"
)

var channel = "#1125"

func main() {

  type Friendlist struct {
    Handle  string
    Host    string
    Flag    string
  }

  type Config struct {
    Nick    string
    Host    string
    Server  string
    Port    string
  }

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

func readConfig() (config byte[]){
  var config Config
  configFile, err := ioutils.ReadFile("config.json")

  if err != nil {
    fmt.Println("Config could not be loaded: " + err)
    return
  }

  err = json.Unmarshal(configFile, &config)

  return config
}
