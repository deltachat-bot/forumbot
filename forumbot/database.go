package main

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"go.uber.org/zap"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

type Post struct {
	Id uint64 `gorm:"primaryKey";autoIncrement:true"`
	Title string
	Body  string
	Author deltachat.ContactId
	CreatedAt time.Time
	Community string
}

type Comment struct {
	Id uint64 `gorm:"primaryKey";autoIncrement:true"`
	Post uint64
	Body  string
	Author deltachat.ContactId
	CreatedAt time.Time
}

func initDB(path string, logger *zap.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Channel{})

	return db
}
