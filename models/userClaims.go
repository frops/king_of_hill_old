package models

import jwt "github.com/dgrijalva/jwt-go"

// UserClaims - is standart pack for user's fields
type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
