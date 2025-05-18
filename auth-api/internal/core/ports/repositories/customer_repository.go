package repositories

import "github.com/PyMarcus/FreelaIF/auth-api/auth-api/internal/core/domain"

type ICustomerRepository interface {
	Create(customer *domain.CustomerEntity) error
	GetCustomerByUserName(username string) (*domain.CustomerEntity, error)
	GetCustomerByEmail(email string) (*domain.CustomerEntity, error)
	ListAll() ([]*domain.CustomerEntity, error)
}
