package dtos

type UserLoginRequestDTO struct{
	Email string 
	Password string
}

type UserRegisterCustomerRequestDTO struct{
	// TODO: adicionar outros campos
	Email string 
	Password string
}

type UserRegisterStudentRequestDTO struct{
	// TODO: adicionar outros campos
	Email string 
	Password string
}