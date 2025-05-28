package tests

import (
	"time"

	"github.com/PyMarcus/FreelaIF/auth-api/auth-api/internal/core/domain"
	"github.com/google/uuid"
)

func createUsersToTest() (*domain.CustomerEntity, *domain.StudentEntity){
	return &domain.CustomerEntity{
			ID:        uuid.New(), 
			Name:      "Empresa Exemplo Ltda",
			Email:     "contato@empresaexemplo.com",
			Username:  "empresa123",
			Password:  "senha123", 
			CNPJ:      "12.345.678/0001-99",
			CPF:       "123.456.789-00",
			Phone:     "(11) 98765-4321",
			Address:   "Rua das Flores, 123, São Paulo, SP",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}, &domain.StudentEntity{
			ID:        uuid.New(),
			Name:      "Joao Silva Foo Bar Baz",
			Email:     "40028922@ifmg.com.br",
			Username:  "joaofoobarbaz",
			Password:  "senha123",
			CPF:       "987.654.321-00",
			Phone:     "(31) 91234-5678",
			Address:   "Av. Brasil, 456, Rio de Janeiro, RJ",
			CreatedAt: time.Now(),
			UpdatedAt: nil,
		}
}

func ptrString(s string) *string {
	return &s
}
