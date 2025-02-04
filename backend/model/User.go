package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
type Session struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token"`
	UserID    string      `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func MigrateSession(db *gorm.DB) {
	// Auto-migrate User and Session tables
	db.AutoMigrate(&User{}, &Session{})
}

// Migrate function to create the table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
