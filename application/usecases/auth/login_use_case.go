package auth_use_cases

import (
	"strings"

	user_use_cases "github.com/karlgama/chat-app-go.git/application/usecases/user"
	"github.com/karlgama/chat-app-go.git/domain/entities"
	postgre "github.com/karlgama/chat-app-go.git/infra/postgreSQL/repositories"
	security "github.com/karlgama/chat-app-go.git/infra/security/services"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginOutput struct {
	Token *string `json:"token"`
}

func Login(input *LoginInput) (*LoginOutput, error) {
	// Usa o repositório real ao invés da função mockada
	repository := &postgre.UserPostgresRepository{}
	foundUser, err := repository.FindUserByEmail(input.Email)

	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return nil, entities.ErrInvalidCredentials
		}
		return nil, err
	}

	if foundUser == nil {
		return nil, entities.ErrInvalidCredentials
	}

	// Verifica se a senha está correta
	isPasswordValid := security.CheckPasswordHash(input.Password, foundUser.Password)
	if !isPasswordValid {
		return nil, entities.ErrInvalidCredentials
	}

	// Gera o token JWT
	token, err := security.GenerateToken(foundUser)
	if err != nil {
		return nil, entities.ErrTokenGeneration
	}

	return &LoginOutput{Token: &token}, nil
}

// Versão estruturada do use case
type LoginUseCase struct {
	userRepository user_use_cases.FindUserByEmailUseCase
}

func (l *LoginUseCase) Execute(input *LoginInput) (*LoginOutput, error) {
	foundUser, err := l.userRepository.Execute(input.Email)

	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return nil, entities.ErrInvalidCredentials
		}
		return nil, err
	}

	if foundUser == nil {
		return nil, entities.ErrInvalidCredentials
	}

	// Verifica se a senha está correta
	isPasswordValid := security.CheckPasswordHash(input.Password, foundUser.Password)
	if !isPasswordValid {
		return nil, entities.ErrInvalidCredentials
	}

	// Gera o token JWT
	token, err := security.GenerateToken(foundUser)
	if err != nil {
		return nil, entities.ErrTokenGeneration
	}

	return &LoginOutput{Token: &token}, nil
}

func NewLoginUseCase(userRepository user_use_cases.FindUserByEmailUseCase) *LoginUseCase {
	return &LoginUseCase{userRepository: userRepository}
}
