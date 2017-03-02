package gate

import (
	"leafchat/server/game"
	"leafchat/server/login"
	"leafchat/server/msg"
)

func init() {
	// login
	msg.JSONProcessor.SetRouter(&msg.C2S_Auth{}, login.ChanRPC)

	// game
	msg.JSONProcessor.SetRouter(&msg.C2S_AddUser{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_Message{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.C2S_Action{}, game.ChanRPC)
}
