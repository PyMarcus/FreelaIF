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

var ErrStudentNotFound = errors.New("student not found")

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
	Students, err := c.usecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var response []dtos.StudentResponseDTO
	for _, Student := range Students {
		response = append(response, mapToStudentResponseDTO(Student))
	}

	ctx.JSON(http.StatusOK, response)
}

// GetStudentByEmail implements httpport.StudentHandler.
func (c StudentHandler) GetStudentByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	Student, err := c.usecase.GetStudentByEmail(email)
	if err != nil {
		if errors.Is(err, ErrStudentNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, mapToStudentResponseDTO(Student))
}

// GetStudentByUserName implements httpport.StudentHandler.
func (c StudentHandler) GetStudentByUserName(ctx *gin.Context) {
	username := ctx.Param("username")

	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	Student, err := c.usecase.GetStudentByUserName(username)
	if err != nil {
		if errors.Is(err, ErrStudentNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, mapToStudentResponseDTO(Student))
}

// Register implements httpport.StudentHandler.
func (c StudentHandler) Register(ctx *gin.Context) {
	var dto dtos.UserRegisterStudentRequestDTO

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	err := c.usecase.Create(&dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Student registered"})
}

// mapToStudentResponseDTO is a helper to convert domain.StudentEntity to dtos.StudentResponseDTO
func mapToStudentResponseDTO(Student *domain.StudentEntity) dtos.StudentResponseDTO {
	return dtos.StudentResponseDTO{
		ID:        Student.ID.String(),
		Name:      Student.Name,
		Email:     Student.Email,
		Username:  Student.Username,
		CPF:       Student.CPF,
		Phone:     Student.Phone,
		Address:   Student.Address,
		CreatedAt: Student.CreatedAt,
	}
}
