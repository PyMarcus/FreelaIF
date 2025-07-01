package dtos

type UserRegisterCustomerRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	CNPJ     string `json:"cnpj,omitempty"`                     // opcional
	CPF      string `json:"cpf" binding:"required,len=11"`      // ou len=14 se considerar pontuação
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UserRegisterStudentRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	CPF      string `json:"cpf" binding:"required,len=11"`      // ou len=14 se considerar pontuação
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}