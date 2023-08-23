package usecase

import (
	"fmt"

	"github.com/NursiNursi/laundry-apps/utils/security"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	usecase UserUseCase
	// JWT
	// service token jwt
}

// Login implements AuthUseCase.
func (a *authUseCase) Login(username string, password string) (string, error) {
	user, err := a.usecase.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	// mekanisme jika user itu ada dan valid dia akan membalikan sebuah token (token123)
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func NewAuthUseCase(usecase UserUseCase) AuthUseCase {
	return &authUseCase{usecase: usecase}
}
