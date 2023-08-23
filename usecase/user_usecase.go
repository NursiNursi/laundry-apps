package usecase

import (
	"fmt"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(paylaod model.UserCredential) error
	FindAllUser() ([]model.UserCredential, error)
	FindByUsername(username string) (model.UserCredential, error)
	FindByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.UserCredential, error) {
	return u.repo.List()
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.repo.GetUsername(username)
}

// FindByUsernamePassword implements UserUseCase.
func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.UserCredential, error) {
	return u.repo.GetUsernamePassword(username, password)
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(paylaod model.UserCredential) error {
	// bytes => sjiadbafiaf7asf8af8as8fasnfajfcnas!dcscsjc
	bytes, _ := bcrypt.GenerateFromPassword([]byte(paylaod.Password), bcrypt.DefaultCost)
	paylaod.Password = string(bytes)
	err := u.repo.Create(paylaod)
	if err != nil {
		return fmt.Errorf("failed to create user %v", err)
	}
	return nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
