package main

import (
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/stretchr/testify/assert"
)

func TestBot(t *testing.T) {
	acfactory.WithOnlineBot(func(bot *deltachat.Bot, botAcc deltachat.AccountId) {
		acfactory.WithOnlineAccount(func(userRpc *deltachat.Rpc, userAcc deltachat.AccountId) {
			bot.OnNewMsg(onNewMsg)
			go bot.Run() //nolint:errcheck

			chatWithBot := acfactory.CreateChat(userRpc, userAcc, bot.Rpc, botAcc)

			_, err := userRpc.MiscSendTextMessage(userAcc, chatWithBot, "hi")
			assert.Nil(t, err)

			msg := acfactory.NextMsg(userRpc, userAcc)
			assert.NotEmpty(t, msg.File)
		})
	})
}
