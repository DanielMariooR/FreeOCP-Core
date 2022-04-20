package models

import "github.com/golang-jwt/jwt"

type SignInInput struct {
	Email    string `json:"email" validate:"required" label:"email"`
	Password string `json:"password" validate:"required" label:"password"`
}

type SignInResponse struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expiration"`
}

type Claims struct {
	ID       string `json:"id"`
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type SignUpInput struct {
	Email    string `json:"email" validate:"required" label:"email"`
	Name     string `json:"fullname" validate:"required" label:"fullname"`
	Username string `json:"username" validate:"required" label:"username"`
	Password string `json:"password" validate:"required" label:"password"`
}

type SignUpResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
