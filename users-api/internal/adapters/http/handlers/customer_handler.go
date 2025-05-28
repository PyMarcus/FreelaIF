package handlers

import (
	"fmt"
	"net/http"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

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
	panic("unimplemented")
}

// GetCustomerByEmail implements httpport.CustomerHandler.
func (c CustomerHandler) GetCustomerByEmail(ctx *gin.Context) {
	panic("unimplemented")
}

// GetCustomerByUserName implements httpport.CustomerHandler.
func (c CustomerHandler) GetCustomerByUserName(ctx *gin.Context) {
	panic("unimplemented")
}

// Register implements httpport.CustomerHandler.
func (c CustomerHandler) Register(ctx *gin.Context) {
	var dto dtos.UserRegisterCustomerRequestDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	err := c.usecase.Create(&dto)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Customer registered"})
}

