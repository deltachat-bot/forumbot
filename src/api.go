package main

import (
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat/xdcrpc"
)

const (
	// Invalid data caused an SQL error
	SQLError xdcrpc.ErrorCode = 1
)

// JSON-RPC API available for the frontend app
type API struct {
	rpc    *deltachat.Rpc
	chatId deltachat.ChatId
	accId  deltachat.AccountId
	msgId  deltachat.MsgId
}

// Up-vote the give post
func (self *API) Like(post uint) *xdcrpc.Error {
	like := &Like{User: self.chatId, AccId: self.accId, PostID: post}
	if err := database.Create(like).Error; err != nil {
		return &xdcrpc.Error{Code: SQLError, Message: err.Error()}
	}
	return nil
}

// Creeate a new post
func (self *API) CreatePost(title string, body string, community string) *xdcrpc.Error {
	post := &Post{
		Author:    self.chatId,
		AccId:     self.accId,
		Title:     title,
		Body:      body,
		Community: community,
	}
	if err := database.Create(post).Error; err != nil {
		return &xdcrpc.Error{Code: SQLError, Message: err.Error()}
	}
	return nil
}
