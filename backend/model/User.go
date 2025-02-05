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
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Token     string    `json:"token" gorm:"uniqueIndex;not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`
}

type Report struct {
	ID        string `json:"id" gorm:"id"`
	UserID    string `json:"uid" gorm:"user_id"`
	Longitude string `json:"longitude" gorm:"longitude"`
	Latitude  string `json:"latitude" gorm:"latitude"`
	Place     string `json:"place" gorm:"place"`
}

func MigrateSession(db *gorm.DB) {
	// Auto-migrate User and Session tables
	db.AutoMigrate(&User{}, &Session{})
}

func MigrateReport(db *gorm.DB) {
	// Auto-migrate User and Session tables
	db.AutoMigrate(&User{}, &Report{})
}

// Migrate function to create the table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
