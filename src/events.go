package main

import (
	"os"
	"path/filepath"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat/option"
)

func onEvent(bot *deltachat.Bot, accId deltachat.AccountId, event deltachat.Event) {
	switch ev := event.(type) {
	case deltachat.EventWebxdcStatusUpdate:
		rawUpdates, err := getUpdates(bot.Rpc, accId, ev.MsgId, ev.StatusUpdateSerial-1)
		if err != nil {
			cli.Logger.Error(err)
			return
		}
		if len(rawUpdates) > 0 {
			handleStatusUpdate(bot, accId, ev.MsgId, rawUpdates[0])
		} else {
			cli.Logger.Warnf("[acc=%v] No new status update was found for event: %#v", accId, ev)
		}
	case deltachat.EventSecurejoinInviterProgress:
		if ev.Progress == 1000 {
			cli.Logger.Debugf("[acc=%v] QR scanned by contact with id=%v", accId, ev.ContactId)
			chatId, err := bot.Rpc.CreateChatByContactId(accId, ev.ContactId)
			if err != nil {
				cli.Logger.Error(err)
				return
			}
			sendApp(bot, accId, chatId)
		}
	}
}

func onNewMsg(bot *deltachat.Bot, accId deltachat.AccountId, msgId deltachat.MsgId) {
	msg, err := bot.Rpc.GetMessage(accId, msgId)
	if err != nil {
		cli.Logger.Error(err)
		return
	}

	if !msg.IsBot && msg.FromId > deltachat.ContactLastSpecial && msg.Text != "" {
		chat, err := bot.Rpc.GetBasicChatInfo(accId, msg.ChatId)
		if err != nil {
			cli.Logger.Error(err)
			return
		}
		if chat.ChatType == deltachat.ChatSingle {
			cli.Logger.Debugf("[acc=%v] Got new 1:1 message: %#v", accId, msg)
			sendApp(bot, accId, msg.ChatId)
		}
	}

	if msg.FromId > deltachat.ContactLastSpecial {
		err = bot.Rpc.DeleteMessages(accId, []deltachat.MsgId{msg.Id})
		if err != nil {
			cli.Logger.Error(err)
		}
	}
}

// send the app / UI interace
func sendApp(bot *deltachat.Bot, accId deltachat.AccountId, chatId deltachat.ChatId) {
	// try to resend existing instance
	none := option.None[deltachat.MsgType]()
	msgIds, err := bot.Rpc.GetChatMedia(accId, chatId, deltachat.MsgWebxdc, none, none)
	if err != nil {
		cli.Logger.Error(err)
		return
	}
	for i := len(msgIds) - 1; i >= 0; i-- {
		msgId := msgIds[i]
		msg, err := bot.Rpc.GetMessage(accId, msgId)
		if err != nil {
			cli.Logger.Error(err)
			continue
		}
		if msg.FromId == deltachat.ContactSelf {
			err = bot.Rpc.ResendMessages(accId, []deltachat.MsgId{msgId})
			if err != nil {
				cli.Logger.Error(err)
				break
			}
			return
		} else {
			err = bot.Rpc.DeleteMessages(accId, []deltachat.MsgId{msgId})
			if err != nil {
				cli.Logger.Error(err)
			}
		}
	}

	// no previous instance exists, send app

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		cli.Logger.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	xdcPath := filepath.Join(dir, "app.xdc")
	if err = os.WriteFile(xdcPath, xdcContent, 0666); err != nil {
		cli.Logger.Error(err)
		return
	}

	_, err = bot.Rpc.SendMsg(accId, chatId, deltachat.MsgData{File: xdcPath})
	if err != nil {
		cli.Logger.Error(err)
	}
}
