package repositories

import "github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"

type IStudentRepository interface{
	Create(student *domain.StudentEntity)(error)
	GetStudentByUserName(username string)(*domain.StudentEntity, error)
	GetStudentByEmail(email string)(*domain.StudentEntity, error)
	ListAll()([]*domain.StudentEntity, error)
}
