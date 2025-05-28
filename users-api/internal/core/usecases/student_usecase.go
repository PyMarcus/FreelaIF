package usecases

import (
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
)

type StudentUsecase struct{
	repo repositories.IStudentRepository
}

func NewStudentUsecase(repo repositories.IStudentRepository) *StudentUsecase{
	return &StudentUsecase{
		repo: repo,
	}
}

func (cu *StudentUsecase) Create(StudentDTO *dtos.UserRegisterStudentRequestDTO) error{
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
