package main

import "github.com/deltachat/deltachat-rpc-client-go/deltachat"

type ErrorCode int

const (
	// The method does not exist / is not available
	MethodNotFoud ErrorCode = -32601
	// Invalid JSON was received by the server
	ParseError ErrorCode = -32700
	// The JSON sent is not a valid Request object
	InvalidRequest ErrorCode = -32600
	// Invalid method parameter(s)
	InvalidParams ErrorCode = -32602
	// Invalid data caused an SQL error
	SQLError ErrorCode = 1
)

type MethodName string

// Request sent by the frontend app
type Request[T any] struct {
	Id     string     `json:"id,omitempty"`
	Method MethodName `json:"method,omitempty"`
	Params T          `json:"params,omitempty"`
}

// Response sent by the bot on success
type ResultResponse[T any] struct {
	Id     string `json:"id"`
	Result T      `json:"result"`
}

// Response sent by the bot on errors
type ErrorResponse struct {
	Id    string `json:"id"`
	Error Error  `json:"error,omitempty"`
}

// Error data sent by the bot in ErrorResponse
type Error struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

// JSON-RPC API available for the frontend app
type API struct {
	rpc       *deltachat.Rpc
	chatId    deltachat.ChatId
	accId     deltachat.AccountId
	msgId     deltachat.MsgId
	requestId string
}

// Up-vote the give post
func (self *API) Like(post uint) *Error {
	like := &Like{User: self.chatId, AccId: self.accId, PostID: post}
	if err := database.Create(like).Error; err != nil {
		return &Error{Code: SQLError, Message: err.Error()}
	}
	return nil
}

// Creeate a new post
func (self *API) CreatePost(title string, body string, community string) *Error {
	post := &Post{
		Author:    self.chatId,
		AccId:     self.accId,
		Title:     title,
		Body:      body,
		Community: community,
	}
	if err := database.Create(post).Error; err != nil {
		return &Error{Code: SQLError, Message: err.Error()}
	}
	return nil
}
