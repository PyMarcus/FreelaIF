package repositories

import (
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/utils"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) repositories.IStudentRepository {
	return &StudentRepository{
		db: db,
	}
}

// Create a Student and hash the password
func (c *StudentRepository) Create(Student *domain.StudentEntity) error {
	hash, err := utils.HashPassword(Student.Password)
	if err != nil{
		return err
	}
	Student.Password = hash
	return c.db.Save(Student).Error
}

// GetStudentByEmail returns a Student by email
func (c *StudentRepository) GetStudentByEmail(email string) (*domain.StudentEntity, error) {
	panic("unimplemented")
}

// GetStudentByUserName returns a Student by username
func (c *StudentRepository) GetStudentByUserName(username string) (*domain.StudentEntity, error) {
	panic("unimplemented")
}

// ListAll returns all Students without restrictions
func (c *StudentRepository) ListAll() ([]*domain.StudentEntity, error) {
	panic("unimplemented")
}
