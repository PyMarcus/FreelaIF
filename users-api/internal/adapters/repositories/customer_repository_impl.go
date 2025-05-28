package repositories

import (
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/utils"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
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
	panic("unimplemented")
}

// GetCustomerByUserName returns a customer by username
func (c *CustomerRepository) GetCustomerByUserName(username string) (*domain.CustomerEntity, error) {
	panic("unimplemented")
}

// ListAll returns all customers without restrictions
func (c *CustomerRepository) ListAll() ([]*domain.CustomerEntity, error) {
	panic("unimplemented")
}
