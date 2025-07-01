package httpport

import (
	"github.com/gin-gonic/gin"
)

type StudentHandler interface {
	Register(ctx *gin.Context)
	GetStudentByUserName(ctx *gin.Context)
	GetStudentByEmail(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
