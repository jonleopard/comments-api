package database

import (
	"github.com/jonleopard/comments-api/internal/comment"
	"gorm.io/gorm"
)

// MigrateDB - migrates our databse and creates our comment table
func MigrateDB(db *gorm.DB) error {
	db.AutoMigrate(&comment.Comment{})
	return nil
}
