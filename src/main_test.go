package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

type TestCallback func(bot *deltachat.Bot, botAcc deltachat.AccountId, userRpc *deltachat.Rpc, userAcc deltachat.AccountId)
type WebxdcCallback func(bot *deltachat.Bot, botAcc deltachat.AccountId, userRpc *deltachat.Rpc, userAcc deltachat.AccountId, msg *deltachat.MsgSnapshot)

var acfactory *deltachat.AcFactory

func TestMain(m *testing.M) {
	acfactory = &deltachat.AcFactory{Debug: os.Getenv("TEST_DEBUG") == "1"}
	acfactory.TearUp()
	defer acfactory.TearDown()
	m.Run()
}

func withBotAndUser(callback TestCallback) {
	acfactory.WithOnlineBot(func(bot *deltachat.Bot, botAcc deltachat.AccountId) {
		acfactory.WithOnlineAccount(func(userRpc *deltachat.Rpc, userAcc deltachat.AccountId) {
			onBotInit(bot, acfactory.MkdirTemp())
			go bot.Run() //nolint:errcheck
			callback(bot, botAcc, userRpc, userAcc)
		})
	})
}

// msg is the webxdc message received in the user side
func withWebxdc(callback WebxdcCallback) {
	acfactory.WithOnlineBot(func(bot *deltachat.Bot, botAcc deltachat.AccountId) {
		acfactory.WithOnlineAccount(func(userRpc *deltachat.Rpc, userAcc deltachat.AccountId) {
			onBotInit(bot, acfactory.MkdirTemp())
			go bot.Run() //nolint:errcheck
			chatWithBot := acfactory.CreateChat(userRpc, userAcc, bot.Rpc, botAcc)

			_, err := userRpc.MiscSendTextMessage(userAcc, chatWithBot, "hi")
			if err != nil {
				panic(err)
			}

			msg := acfactory.NextMsg(userRpc, userAcc)
			if !strings.HasSuffix(msg.File, ".xdc") {
				panic("unexpected file name: " + msg.File)
			}

			callback(bot, botAcc, userRpc, userAcc, msg)
		})
	})
}

// Get the Payload contained in the status update with the given serial
func getTestPayload[T any](rpc *deltachat.Rpc, accId deltachat.AccountId, msgId deltachat.MsgId, serial uint) T {
	rawUpdates, err := getUpdates(rpc, accId, msgId, serial-1)
	if err != nil {
		panic(err)
	}
	var update StatusUpdate[T]
	err = json.Unmarshal(rawUpdates[0], &update)
	if err != nil {
		panic(err)
	}
	return update.Payload
}

// Get bot response
func getTestResponse[T any](rpc *deltachat.Rpc, accId deltachat.AccountId) T {
	ev := acfactory.WaitForEvent(rpc, accId, deltachat.EventWebxdcStatusUpdate{}).(deltachat.EventWebxdcStatusUpdate)
	return getTestPayload[T](rpc, accId, ev.MsgId, ev.StatusUpdateSerial)
}

// Send a status update with the given request.
// Automatically ignore the next EventWebxdcStatusUpdate from self
func sendTestRequest[T any](rpc *deltachat.Rpc, accId deltachat.AccountId, msgId deltachat.MsgId, req Request[T]) {
	sendPayload(rpc, accId, msgId, req)

	// ignore self-update
	ev := acfactory.WaitForEvent(rpc, accId, deltachat.EventWebxdcStatusUpdate{}).(deltachat.EventWebxdcStatusUpdate)
	resp := getTestPayload[Request[T]](rpc, accId, ev.MsgId, ev.StatusUpdateSerial)
	if resp.Id != req.Id {
		panic("Unexpected request Id: " + resp.Id)
	}
}
