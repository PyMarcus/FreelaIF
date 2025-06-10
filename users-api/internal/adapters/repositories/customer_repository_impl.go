package repositories

import (
	"time"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/utils"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repositories.ICustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

// Create a customer and hash the password
func (c *CustomerRepository) Create(customer *domain.CustomerEntity) error {
	hash, err := utils.HashPassword(customer.Password)
	if err != nil {
		return err
	}
	customer.Password = hash
	return c.db.Save(customer).Error
}

// GetCustomerByEmail returns a customer by email
func (c *CustomerRepository) GetCustomerByEmail(email string) (*domain.CustomerEntity, error) {
	stubCustomer1 := domain.CustomerEntity{
		ID:        uuid.New(),
		Name:      "João da Silva",
		Email:     "joao.silva@example.com",
		Username:  "joaosilva",
		Password:  "hashedpassword123",
		CNPJ:      "12.345.678/0001-95",
		CPF:       "123.456.789-00",
		Phone:     "+55 11 91234-5678",
		Address:   "Rua das Flores, 123 - São Paulo, SP",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	return &stubCustomer1, nil
}

// GetCustomerByUserName returns a customer by username
func (c *CustomerRepository) GetCustomerByUserName(username string) (*domain.CustomerEntity, error) {
	stubCustomer1 := domain.CustomerEntity{
		ID:        uuid.New(),
		Name:      "João da Silva",
		Email:     "joao.silva@example.com",
		Username:  "joaosilva",
		Password:  "hashedpassword123",
		CNPJ:      "12.345.678/0001-95",
		CPF:       "123.456.789-00",
		Phone:     "+55 11 91234-5678",
		Address:   "Rua das Flores, 123 - São Paulo, SP",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	return &stubCustomer1, nil
}

// ListAll returns all customers without restrictions
func (c *CustomerRepository) ListAll() ([]*domain.CustomerEntity, error) {

	stubCustomer1 := domain.CustomerEntity{
		ID:        uuid.New(),
		Name:      "João da Silva",
		Email:     "joao.silva@example.com",
		Username:  "joaosilva",
		Password:  "hashedpassword123",
		CNPJ:      "12.345.678/0001-95",
		CPF:       "123.456.789-00",
		Phone:     "+55 11 91234-5678",
		Address:   "Rua das Flores, 123 - São Paulo, SP",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}

	stubCustomer2 := domain.CustomerEntity{
		ID:        uuid.New(),
		Name:      "Maria Oliveira",
		Email:     "maria.oliveira@example.com",
		Username:  "mariaoliveira",
		Password:  "hashedpassword456",
		CNPJ:      "98.765.432/0001-12",
		CPF:       "987.654.321-00",
		Phone:     "+55 21 99876-5432",
		Address:   "Av. Atlântica, 456 - Rio de Janeiro, RJ",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}
	return []*domain.CustomerEntity{
		&stubCustomer1,
		&stubCustomer2,
	}, nil
}
