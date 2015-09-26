package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Friendlist struct {
	Handle string
	Host   string
	Flag   string
}

type Config struct {
	Nick    string
	Host    string
	Server  string
	Port    string
	Channel string
}

func main() {

	config := readConfig()
	friendlist := loadFriendlist()
	con := irc.IRC(config.Nick, config.Nick)
	err := con.Connect(config.Server + ":" + config.Port)

	if err != nil {
		fmt.Println("Connection failed")
		return
	}

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(config.Channel)
	})

	con.AddCallback("JOIN", func(e *irc.Event) {
		con.Privmsg(e.Arguments[0], "Sup!")
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		msgParts := strings.Split(e.Arguments[1], " ")

		if friendlist.Host == (e.User + "@" + e.Host) {
			switch {
			case msgParts[0] == "!quit":
				if len(msgParts) > 1 {
					con.SendRaw("QUIT :" + msgParts[1])
				} else {
					con.SendRaw("QUIT :I quit")
				}
				con.Quit()

			case msgParts[0] == "!join":
				con.Join(msgParts[1])

			default:
				fmt.Println(e.Arguments)
			}
		}
	})

	con.Loop()
}

func readConfig() Config {
	var config Config
	configFile, err := ioutil.ReadFile("config")

	if err != nil {
		fmt.Println("Config could not be loaded:")
		fmt.Println(err)
	}

	err = yaml.Unmarshal(configFile, &config)

	return config
}

func loadFriendlist() Friendlist {
	var friendlist Friendlist
	friendlistFile, err := ioutil.ReadFile("friendlist")

	if err != nil {
		fmt.Println("Friendlist could not be loaded")
		fmt.Println(err)
	}

	err = yaml.Unmarshal(friendlistFile, &friendlist)

	return friendlist
}
