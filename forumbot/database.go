package main

import (
	"time"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Post struct {
	Id        uint64 `gorm:"primaryKey";autoIncrement:true"`
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

func initDB(path string, logger *zap.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Post{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	err = db.AutoMigrate(&Like{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return db
}
