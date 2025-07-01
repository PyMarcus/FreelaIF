package repositories

import "github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"

type ICustomerRepository interface {
	Create(customer *domain.CustomerEntity) error
	GetCustomerByUserName(username string) (*domain.CustomerEntity, error)
	GetCustomerByEmail(email string) (*domain.CustomerEntity, error)
	ListAll() ([]*domain.CustomerEntity, error)
}
//mockgen -source=internal/core/ports/repositories/customer_repository.go -destination=tests/mocks/mock_customer_repository.go -package=mocks