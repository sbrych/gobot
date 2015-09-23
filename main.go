package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var channel = "#1125"

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

func main() {

  config  := readConfig()
	con     := irc.IRC(config.Nick, config.Nick)
  err     := con.Connect(config.Server + ":" + config.Port)

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
