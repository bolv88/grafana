package models

import "errors"

var ErrInvalidEmailCode = errors.New("Invalid or expired email code")

type SendEmailCommand struct {
	To       []string
	Template string
	Data     map[string]interface{}
	Massive  bool
	Info     string
}

type SendWebhook struct {
	Url          string
	AuthUser     string
	AuthPassword string
	Body         string
	Method       string
}

type SendResetPasswordEmailCommand struct {
	User *User
}

type ValidateResetPasswordCodeQuery struct {
	Code   string
	Result *User
}