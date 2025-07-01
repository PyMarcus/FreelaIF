package usecases

import (
	"errors"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
)

type CustomerUsecase struct {
	repo repositories.ICustomerRepository
}

func NewCustomerUsecase(repo repositories.ICustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{
		repo: repo,
	}
}

func (cu *CustomerUsecase) Create(customerDTO *dtos.UserRegisterCustomerRequestDTO) error {
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

func (cu *CustomerUsecase) GetAll() ([]*domain.CustomerEntity, error) {
	return cu.repo.ListAll()
}

func (cu *CustomerUsecase) GetCustomerByEmail(email string) (*domain.CustomerEntity, error) {
	customer, err := cu.repo.GetCustomerByEmail(email)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (cu *CustomerUsecase) GetCustomerByUserName(username string) (*domain.CustomerEntity, error) {
	customer, err := cu.repo.GetCustomerByUserName(username)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}
