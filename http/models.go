package models

import "time"

type CommonResponse struct {
	Code      int         `json:"code"`
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type AuthInfo struct {
	Status  int          `json:"status"`
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    AuthInfoData `json:"data"`
}

type AuthInfoData struct {
	Id              int64       `json:"id"`
	Username        string      `json:"username"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Email           string      `json:"email"`
	EmailVerifiedAt interface{} `json:"email_verified_at"`
	PhotoPath       string      `json:"photo_path"`
	Status          int         `json:"status"`
	RoleId          int         `json:"role_id"`
	LastLoginAt     string      `json:"last_login_at"`
	LastLoginIp     string      `json:"last_login_ip"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Name            interface{} `json:"name"`
}
