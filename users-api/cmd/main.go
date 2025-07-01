package main

import (
	"log"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/config"
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/http/handlers"
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/http/routes"
	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/repositories"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/usecases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	registerInEurekaServer()
	go heartBeat()

	connStr := config.ConfigSettings.DatabaseUrl
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	customerRepository := repositories.NewCustomerRepository(db)
	customerUsecase := usecases.NewCustomerUsecase(
		customerRepository,
	)
	customerHandler := handlers.NewCustomerHandler(customerUsecase)

	studentRepository := repositories.NewStudentRepository(db)
	studentUsecase := usecases.NewStudentUsecase(
		studentRepository,
	)
	studentHandler := handlers.NewStudentHandler(studentUsecase)

	authRepository := repositories.NewAuthRepository(db)
	authUsecase := usecases.NewAuthUsecase(authRepository)
	authHandler := handlers.NewLoginHandler(authUsecase)

	router := routes.SetupRouter(customerHandler, studentHandler, authHandler)

	if err := router.Run(":8080"); err != nil{
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

}