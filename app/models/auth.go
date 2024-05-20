package models

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Register
type RegisterRequest struct {
	Username *string `json:"username" validate:"required,min=1"`
	Password string  `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Login
type LoginRequest struct {
	Username *string `json:"username" validate:"required,min=1"`
	Password string  `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
