package usecases

import (
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
)

type CustomerUsecase struct{
	repo repositories.ICustomerRepository
}

func NewCustomerUsecase(repo repositories.ICustomerRepository) *CustomerUsecase{
	return &CustomerUsecase{
		repo: repo,
	}
}

func (cu *CustomerUsecase) Create(customerDTO *dtos.UserRegisterCustomerRequestDTO) error{
	customer := &domain.CustomerEntity{
		Name:     customerDTO.Name,
		Username: customerDTO.Username,
		Email:    customerDTO.Email,
		Password: customerDTO.Password,
		CNPJ:     customerDTO.CNPJ,
		CPF:      customerDTO.CPF,
		Phone:    customerDTO.Phone,
		Address:  customerDTO.Address,
	}
	return cu.repo.Create(customer)
}
