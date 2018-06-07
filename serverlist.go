package main

import (
	"os"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
	"github.com/MarcPlaying/go-serverlist/constants"
	"github.com/MarcPlaying/go-serverlist/server"
)

func init() {
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		conf := &constants.Config{}
		conf.MySQL = consts.MySQLConf{
			Database: "serverlist",
			Hostname: "0.0.0.0",
			Port:     3306,
			Username: "root",
		}
		conf.Server = struct {
			Host  string
			Port  int16
			Debug bool
		}{
			Host: "0.0.0.0",
			Port: 8745,
		}
		conf.Redis = consts.REDISConf{
			Host: "127.0.0.1",
			Port: 6379,
		}
		helpers.CreateConfig("config", conf)
		logger.Info("I've just created a config for you! Please edit.")
		os.Exit(0)
	}
}

func main() {
	server.StartServer("", 8745)
}
