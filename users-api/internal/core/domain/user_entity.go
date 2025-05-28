package domain

import (
	"time"

	"github.com/google/uuid"
)

type StudentEntity struct{
 	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name      string
    Email     string `gorm:"uniqueIndex"`
    Username  string `gorm:"uniqueIndex"`
    Password  string
    CPF       string `gorm:"uniqueIndex"`
    Phone     string
    Address   string
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt *time.Time
}

type CustomerEntity struct{
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name      string
    Email     string `gorm:"uniqueIndex"`
    Username  string `gorm:"uniqueIndex"`
    Password  string
    CNPJ      string
    CPF       string `gorm:"uniqueIndex"`
    Phone     string
    Address   string
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt *time.Time
}

func (CustomerEntity) TableName() string{
    return "customers"
}
