package usecases

import (
	"errors"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
)

type StudentUsecase struct {
	repo repositories.IStudentRepository
}

func NewStudentUsecase(repo repositories.IStudentRepository) *StudentUsecase {
	return &StudentUsecase{
		repo: repo,
	}
}

func (cu *StudentUsecase) Create(StudentDTO *dtos.UserRegisterStudentRequestDTO) error {
	Student := &domain.StudentEntity{
		Name:     StudentDTO.Name,
		Username: StudentDTO.Username,
		Email:    StudentDTO.Email,
		Password: StudentDTO.Password,
		CPF:      StudentDTO.CPF,
		Phone:    StudentDTO.Phone,
		Address:  StudentDTO.Address,
	}
	return cu.repo.Create(Student)
}

func (cu *StudentUsecase) GetAll() ([]*domain.StudentEntity, error) {
	return cu.repo.ListAll()
}

func (cu *StudentUsecase) GetStudentByEmail(email string) (*domain.StudentEntity, error) {
	Student, err := cu.repo.GetStudentByEmail(email)
	if err != nil {
		return nil, err
	}
	if Student == nil {
		return nil, errors.New("student not found")
	}
	return Student, nil
}

func (cu *StudentUsecase) GetStudentByUserName(username string) (*domain.StudentEntity, error) {
	Student, err := cu.repo.GetStudentByUserName(username)
	if err != nil {
		return nil, err
	}
	if Student == nil {
		return nil, errors.New("student not found")
	}
	return Student, nil
}
