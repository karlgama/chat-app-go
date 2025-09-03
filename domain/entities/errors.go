package entities

import "errors"

// Erros de domínio para autenticação
var (
	ErrInvalidCredentials = errors.New("email or password is incorrect")
	ErrUserNotFound       = errors.New("user not found")
	ErrTokenGeneration    = errors.New("could not generate token")
)

// Erros de domínio para usuários
var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserData   = errors.New("invalid user data")
)

// Erros de domínio para chats
var (
	ErrChatNotFound    = errors.New("chat not found")
	ErrInvalidChatData = errors.New("invalid chat data")
)
