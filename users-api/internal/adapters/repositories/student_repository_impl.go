package repositories

import (
	"time"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/utils"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
	"gorm.io/gorm"
	"github.com/google/uuid"
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
	var stubStudent1 = domain.StudentEntity{
			ID:        uuid.New(),
			Name:      "Lucas Martins",
			Email:     "lucas.martins@student.com",
			Username:  "lucasm",
			Password:  "securehash1",
			CPF:       "321.654.987-00",
			Phone:     "+55 31 98888-1111",
			Address:   "Rua Estudante Feliz, 100 - Belo Horizonte, MG",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}
	return &stubStudent1, nil
}

// GetStudentByUserName returns a Student by username
func (c *StudentRepository) GetStudentByUserName(username string) (*domain.StudentEntity, error) {
		var stubStudent1 = domain.StudentEntity{
			ID:        uuid.New(),
			Name:      "Lucas Martins",
			Email:     "lucas.martins@student.com",
			Username:  "lucasm",
			Password:  "securehash1",
			CPF:       "321.654.987-00",
			Phone:     "+55 31 98888-1111",
			Address:   "Rua Estudante Feliz, 100 - Belo Horizonte, MG",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}
	return &stubStudent1, nil
}

// ListAll returns all Students without restrictions
func (c *StudentRepository) ListAll() ([]*domain.StudentEntity, error) {
	var stubStudent1 = domain.StudentEntity{
	ID:        uuid.New(),
	Name:      "Lucas Martins",
	Email:     "lucas.martins@student.com",
	Username:  "lucasm",
	Password:  "securehash1",
	CPF:       "321.654.987-00",
	Phone:     "+55 31 98888-1111",
	Address:   "Rua Estudante Feliz, 100 - Belo Horizonte, MG",
	CreatedAt: time.Now(),
	UpdatedAt: nil,
}

var stubStudent2 = domain.StudentEntity{
	ID:        uuid.New(),
	Name:      "Ana Paula Souza",
	Email:     "ana.souza@student.com",
	Username:  "anapsouza",
	Password:  "securehash2",
	CPF:       "654.321.789-00",
	Phone:     "+55 41 97777-2222",
	Address:   "Av. do Saber, 200 - Curitiba, PR",
	CreatedAt: time.Now(),
	}

	return []*domain.StudentEntity{&stubStudent1, &stubStudent2}, nil
}
