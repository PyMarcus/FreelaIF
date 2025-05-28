package routes

import (
	httpport "github.com/PyMarcus/FreelaIF/auth-api/auth-api/internal/core/ports/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(customerHandler httpport.CustomerHandler) *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/customers", customerHandler.Register)
	}

	return router
}

