package dtos

import "time"

type CustomerResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CPF       string    `json:"cpf"`
	CNPJ      string    `json:"cnpj,omitempty"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" `
}

type StudentResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CPF       string    `json:"cpf"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterCustomerResponse struct {
	Message  string              `json:"message"`
	Customer CustomerResponseDTO `json:"customer"`
}

type RegisterStudentResponse struct {
	Message string             `json:"message"`
	Student StudentResponseDTO `json:"student"`
}