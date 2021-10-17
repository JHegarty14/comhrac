package users

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

// User => type definition for the User model
type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Age       int
	Height    string
	Weight    int
	Email     string `gorm:"index"`
	Zipcode   string
	CountryID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// UserSettings => type definition for the UserSettings model
type UserSettings struct {
	ID                  int `gorm:"primaryKey"`
	Display             string
	Voice               string
	SportPreference     string
	SyncWithAppleHealth bool
	Stance              string
}

// Token => type definition for the Token model
type Token struct {
	UserID         int `gorm:"primaryKey"`
	StandardClaims *jwt.StandardClaims
}
