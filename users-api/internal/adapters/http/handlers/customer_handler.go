package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/domain"
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

var ErrCustomerNotFound = errors.New("customer not found")

type CustomerHandler struct {
	usecase *usecases.CustomerUsecase
}

func NewCustomerHandler(usecase *usecases.CustomerUsecase) httpport.CustomerHandler {
	return CustomerHandler{
		usecase: usecase,
	}
}

// GetAll implements httpport.CustomerHandler.
func (c CustomerHandler) GetAll(ctx *gin.Context) {
	customers, err := c.usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var response []dtos.CustomerResponseDTO
	for _, customer := range customers {
		response = append(response, mapToCustomerResponseDTO(customer))
	}

	ctx.JSON(http.StatusOK, response)
}

// GetCustomerByEmail implements httpport.CustomerHandler.
func (c CustomerHandler) GetCustomerByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	customer, err := c.usecase.GetCustomerByEmail(email)
	if err != nil {
		if errors.Is(err, ErrCustomerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, mapToCustomerResponseDTO(customer))
}

// GetCustomerByUserName implements httpport.CustomerHandler.
func (c CustomerHandler) GetCustomerByUserName(ctx *gin.Context) {
	username := ctx.Param("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	customer, err := c.usecase.GetCustomerByUserName(username)
	if err != nil {
		if errors.Is(err, ErrCustomerNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, mapToCustomerResponseDTO(customer))
}

// Register implements httpport.CustomerHandler.
func (c CustomerHandler) Register(ctx *gin.Context) {
	var dto dtos.UserRegisterCustomerRequestDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	err := c.usecase.Create(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Customer registered"})
}

// mapToCustomerResponseDTO is a helper to convert domain.CustomerEntity to dtos.CustomerResponseDTO
func mapToCustomerResponseDTO(customer *domain.CustomerEntity) dtos.CustomerResponseDTO {
	return dtos.CustomerResponseDTO{
		ID:        customer.ID.String(),
		Name:      customer.Name,
		Email:     customer.Email,
		Username:  customer.Username,
		CPF:       customer.CPF,
		CNPJ:      customer.CNPJ,
		Phone:     customer.Phone,
		Address:   customer.Address,
		CreatedAt: customer.CreatedAt,
	}
}
