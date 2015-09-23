package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
)

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
  Channel string
}

func main() {

  config  := readConfig()
	con     := irc.IRC(config.Nick, config.Nick)
  err     := con.Connect(config.Server + ":" + config.Port)

	if err != nil {
		fmt.Println("Connection failed")
		return
	}

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(config.Channel)
	})

	con.AddCallback("JOIN", func(e *irc.Event) {
		con.Privmsg(config.Channel, "Sup!")
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
    msgParts := strings.Split(e.Arguments[1], " ")

		switch {
    case msgParts[0] == "!quit":
      if msgParts[1] == "" {
        con.SendRaw("QUIT :I quit")
      } else {
        con.SendRaw("QUIT :" + msgParts[1])
      }
      con.Quit()

    case msgParts[0] == "!join":
      con.Join(msgParts[1])

    default:
      fmt.Println(e.Arguments)
    }
	})

	con.Loop()
}

func readConfig() (Config) {
  var config Config
  configFile, err := ioutil.ReadFile("config.yml")

  if err != nil {
    fmt.Println("Config could not be loaded:")
    fmt.Println(err)
  }

  err = yaml.Unmarshal(configFile, &config)

  return config
}
