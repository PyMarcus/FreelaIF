package usecases

import (
		"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
			r "github.com/PyMarcus/FreelaIF/users-api/internal/adapters/repositories"

)

type AuthUsecase struct {
	repo *repositories.IAuthRepository[r.User]
}

func NewAuthUsecase(repo *repositories.IAuthRepository[r.User]) repositories.IAuthRepository[r.User] {
	return &AuthUsecase{
		repo: repo,
	}
}

func (au AuthUsecase) GetUser(email, password string) (r.User, error) {
	user, _ := au.GetUser(email, password)
	return (user, nil)
}