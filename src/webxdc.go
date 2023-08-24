package main

import (
	"encoding/json"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

var (
	InvalidParamsErr = &Error{Code: InvalidParams, Message: "Invalid params"}
	MethodNotFoudErr = &Error{Code: MethodNotFoud, Message: "Method not found"}
)

const MethodLike MethodName = "Like"

type LikeParams struct {
	Post uint
}

const MethodCreatePost MethodName = "CreatePost"

type CreatePostParams struct {
	Title     string
	Body      string
	Community string
}

type StatusUpdate[T any] struct {
	Info      string `json:"info,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Document  string `json:"document,omitempty"`
	Payload   T      `json:"payload,omitempty"`
	Serial    uint   `json:"serial,omitempty"`
	MaxSerial uint   `json:"max_serial,omitempty"`
}

// handle a webxdc status update
func handleStatusUpdate(bot *deltachat.Bot, accId deltachat.AccountId, msgId deltachat.MsgId, rawUpdate json.RawMessage) {
	if isFromSelf(rawUpdate) {
		cli.Logger.Debugf("[acc=%v] [WebXDC] Response: %v", accId, string(rawUpdate))
		return
	}

	msg, err := bot.Rpc.GetMessage(accId, msgId)
	if err != nil {
		cli.Logger.Error(err)
		return
	}
	chat, err := bot.Rpc.GetBasicChatInfo(accId, msg.ChatId)
	if err != nil {
		cli.Logger.Error(err)
		return
	}
	if chat.ChatType != deltachat.ChatSingle {
		cli.Logger.Debugf("[acc=%v] [WebXDC] Ignoring request in multi-user chat #%v: %v", accId, chat.Id, string(rawUpdate))
		return
	}

	var update StatusUpdate[Request[json.RawMessage]]
	err = json.Unmarshal(rawUpdate, &update)
	if err != nil {
		cli.Logger.Debugf("[acc=%v] [WebXDC] Invalid update: %v", accId, string(rawUpdate))
		return
	}

	cli.Logger.Debugf("[acc=%v] [WebXDC] Request: %v", accId, string(rawUpdate))

	request := update.Payload
	api := &API{rpc: bot.Rpc, chatId: chat.Id, accId: accId, msgId: msgId, requestId: request.Id}
	switch request.Method {
	case MethodLike:
		var data LikeParams
		err := json.Unmarshal(request.Params, &data)
		if err != nil {
			sendResponse[any](api, nil, InvalidParamsErr)
			return
		}
		sendResponse[any](api, nil, api.Like(data.Post))
	case MethodCreatePost:
		var data CreatePostParams
		err := json.Unmarshal(request.Params, &data)
		if err != nil {
			sendResponse[any](api, nil, InvalidParamsErr)
			return
		}
		sendResponse[any](api, nil, api.CreatePost(data.Title, data.Body, data.Community))
	default:
		sendResponse[any](api, nil, MethodNotFoudErr)
	}
}

// Get all setatus updates with serial greater than the given serial
func getUpdates(rpc *deltachat.Rpc, accId deltachat.AccountId, msgId deltachat.MsgId, serial uint) ([]json.RawMessage, error) {
	var rawUpdates []json.RawMessage
	data, err := rpc.GetWebxdcStatusUpdates(accId, msgId, serial)
	if err != nil {
		return rawUpdates, err
	}
	err = json.Unmarshal([]byte(data), &rawUpdates)
	return rawUpdates, err
}

// Send a WebXDC status update
func sendUpdate[T any](rpc *deltachat.Rpc, accId deltachat.AccountId, msgId deltachat.MsgId, update T) {
	data, err := json.Marshal(update)
	if err != nil {
		cli.Logger.Error(err)
		return
	}
	err = rpc.SendWebxdcStatusUpdate(accId, msgId, string(data), "")
	if err != nil {
		cli.Logger.Error(err)
	}
}

// Send a WebXDC status update with the given payload
func sendPayload[T any](rpc *deltachat.Rpc, accId deltachat.AccountId, msgId deltachat.MsgId, payload T) {
	sendUpdate(rpc, accId, msgId, StatusUpdate[T]{Payload: payload})
}

func isFromSelf(rawUpdate json.RawMessage) bool {
	var update StatusUpdate[map[string]json.RawMessage]
	if err := json.Unmarshal(rawUpdate, &update); err != nil {
		return false
	}
	if _, ok := update.Payload["result"]; ok {
		return true
	}
	if _, ok := update.Payload["error"]; ok {
		return true
	}
	return false
}

func sendResponse[T any](api *API, result T, err *Error) {
	if api.requestId == "" {
		return
	}
	if err != nil {
		response := ErrorResponse{Id: api.requestId, Error: *err}
		sendPayload(api.rpc, api.accId, api.msgId, response)
	} else {
		response := ResultResponse[T]{Id: api.requestId, Result: result}
		sendPayload(api.rpc, api.accId, api.msgId, response)
	}
}
