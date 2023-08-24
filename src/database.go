package main

import (
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

type Post struct {
	gorm.Model
	Author    deltachat.ChatId
	AccId     deltachat.AccountId
	Title     string
	Body      string
	InReplyTo uint
	Thread    uint
	Community string
}

type Like struct {
	User   deltachat.ChatId    `gorm:"primaryKey"`
	AccId  deltachat.AccountId `gorm:"primaryKey"`
	PostID uint
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func initDB(path string) {
	var err error
	database, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		cli.Logger.Error(err.Error())
		panic("failed to connect database")
	}

	if res := database.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		cli.Logger.Error(res.Error.Error())
		panic(res.Error)
	}

	// Migrate the schema
	if err := database.AutoMigrate(&Post{}); err != nil {
		cli.Logger.Error(err.Error())
		panic(err)
	}
	if err := database.AutoMigrate(&Like{}); err != nil {
		cli.Logger.Error(err.Error())
		panic(err)
	}
}
