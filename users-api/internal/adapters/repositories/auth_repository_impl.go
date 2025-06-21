package repositories

import (
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/repositories"
	"gorm.io/gorm"
)

type AuthRepository struct{
	db *gorm.DB
}

type User struct{
	ID int 
	Email string 
	Name string
}

func NewAuthRepository(db *gorm.DB) repositories.IAuthRepository[User]{
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) GetUser(email string, password string) (User, error){
	userDump := User{ID: 1, Name: "Lucas Martins", Email: "lucas.martins@student.com"}
	return userDump, nil 
}
