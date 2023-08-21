package main

import (
	"time"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

type Post struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Author    deltachat.ContactId
	Title     string
	Body      string
	CreatedAt time.Time
	InReplyTo uint64
	Thread    uint64
	Community string
}

type Like struct {
	User deltachat.ContactId
	Post uint64
}

func initDB(path string) {
	var err error
	database, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		cli.Logger.Error(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	err = database.AutoMigrate(&Post{})
	if err != nil {
		cli.Logger.Error(err.Error())
		panic(err)
	}
	err = database.AutoMigrate(&Like{})
	if err != nil {
		cli.Logger.Error(err.Error())
		panic(err)
	}
}
