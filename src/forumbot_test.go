package main

import (
	"strings"
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/stretchr/testify/assert"
)

func TestOnNewMsg(t *testing.T) {
	withBotAndUser(func(bot *deltachat.Bot, botAcc deltachat.AccountId, userRpc *deltachat.Rpc, userAcc deltachat.AccountId) {
		chatWithBot := acfactory.CreateChat(userRpc, userAcc, bot.Rpc, botAcc)

		_, err := userRpc.MiscSendTextMessage(userAcc, chatWithBot, "hi")
		assert.Nil(t, err)

		msg := acfactory.NextMsg(userRpc, userAcc)
		assert.True(t, strings.HasSuffix(msg.File, ".xdc"))
	})
}

func TestWebxdcInvalidRequest(t *testing.T) {
	withWebxdc(func(bot *deltachat.Bot, botAcc deltachat.AccountId, userRpc *deltachat.Rpc, userAcc deltachat.AccountId, msg *deltachat.MsgSnapshot) {
		req := Request[any]{Id: "req1", Method: "invalidMethod"}
		sendTestRequest(userRpc, userAcc, msg.Id, req)
		resp := getTestResponse[ErrorResponse](userRpc, userAcc)
		assert.Equal(t, req.Id, resp.Id)
		assert.Equal(t, MethodNotFoud, resp.Error.Code)

		req = Request[any]{Id: "req2", Method: MethodLike, Params: LikeParams{Post: 1}}
		sendTestRequest(userRpc, userAcc, msg.Id, req)
		resp = getTestResponse[ErrorResponse](userRpc, userAcc)
		assert.Equal(t, req.Id, resp.Id)
		assert.Equal(t, SQLError, resp.Error.Code)
	})
}

func TestWebxdc(t *testing.T) {
	withWebxdc(func(bot *deltachat.Bot, botAcc deltachat.AccountId, userRpc *deltachat.Rpc, userAcc deltachat.AccountId, msg *deltachat.MsgSnapshot) {
		params := CreatePostParams{Title: "Test Post", Body: "Test Body", Community: "en"}
		req := Request[any]{Id: "req1", Method: MethodCreatePost, Params: params}
		sendTestRequest(userRpc, userAcc, msg.Id, req)
		resp := getTestResponse[map[string]any](userRpc, userAcc)
		assert.Contains(t, resp, "id")
		assert.Equal(t, req.Id, resp["id"])
		assert.Contains(t, resp, "result")

		req = Request[any]{Id: "req2", Method: MethodLike, Params: LikeParams{Post: 1}}
		sendTestRequest(userRpc, userAcc, msg.Id, req)
		resp = getTestResponse[map[string]any](userRpc, userAcc)
		assert.Contains(t, resp, "id")
		assert.Equal(t, req.Id, resp["id"])
		assert.Contains(t, resp, "result")
		assert.Nil(t, resp["result"])
	})
}
