package main

import (
	"leafchat/server/conf"
	"leafchat/server/game"
	"leafchat/server/gate"
	"leafchat/server/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
