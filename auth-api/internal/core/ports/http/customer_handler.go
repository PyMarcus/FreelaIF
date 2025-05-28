package httpport

import (
	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	Register(ctx *gin.Context)
	GetCustomerByUserName(ctx *gin.Context)
	GetCustomerByEmail(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
