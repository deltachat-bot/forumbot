package main

import (
	"encoding/json"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

type StatusUpdate[T any] struct {
	Type string
	Data *T
}

// handle a webxdc status update
func handleStatusUpdate(bot *deltachat.Bot, accId deltachat.AccountId, msgId deltachat.MsgId, rawUpdate json.RawMessage) {
	var update StatusUpdate[string]
	err := json.Unmarshal(rawUpdate, update)
	if err != nil {
		cli.Logger.Error(err)
		return
	}

	// todo
	err = bot.Rpc.SendWebxdcStatusUpdate(accId, msgId, "TODO: update", "")
	if err != nil {
		cli.Logger.Error(err)
	}
}
