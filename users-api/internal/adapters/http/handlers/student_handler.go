package handlers

import (
	"fmt"
	"net/http"

	"github.com/PyMarcus/FreelaIF/users-api/internal/adapters/dtos"
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/PyMarcus/FreelaIF/users-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	usecase *usecases.StudentUsecase
}

func NewStudentHandler(usecase *usecases.StudentUsecase) httpport.StudentHandler {
	return StudentHandler{
		usecase: usecase,
	}
}


// GetAll implements httpport.StudentHandler.
func (c StudentHandler) GetAll(ctx *gin.Context) {
	panic("unimplemented")
}

// GetStudentByEmail implements httpport.StudentHandler.
func (c StudentHandler) GetStudentByEmail(ctx *gin.Context) {
	panic("unimplemented")
}

// GetStudentByUserName implements httpport.StudentHandler.
func (c StudentHandler) GetStudentByUserName(ctx *gin.Context) {
	panic("unimplemented")
}

// Register implements httpport.StudentHandler.
func (c StudentHandler) Register(ctx *gin.Context) {
	var dto dtos.UserRegisterStudentRequestDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	err := c.usecase.Create(&dto)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Student registered"})
}

