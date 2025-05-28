package routes

import (
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(customerHandler httpport.CustomerHandler, studentHandler httpport.StudentHandler) *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/customers", customerHandler.Register)
		v1.POST("/students", studentHandler.Register)
	}

	return router
}

