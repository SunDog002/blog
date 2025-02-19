package model

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int    `gorm:"type:int;unsigned" json:"id"`
	Username string `gorm:"type:varchar(32);not null" json:"username"`
	Password string `gorm:"type:char(32);not null" json:"password"`
}

type UserClaims struct {
	Username string
	jwt.RegisteredClaims
}
