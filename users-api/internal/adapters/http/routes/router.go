package routes

import (
	httpport "github.com/PyMarcus/FreelaIF/users-api/internal/core/ports/http"
	"github.com/gin-gonic/gin"
)

func SetupRouter(customerHandler httpport.CustomerHandler, studentHandler httpport.StudentHandler, authHandler httpport.AuthHandler) *gin.Engine{
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Customer endpoints
		v1.POST("/customers", customerHandler.Register)
		v1.GET("/customers", customerHandler.GetAll)
		v1.GET("/customers/by-email/:email", customerHandler.GetCustomerByEmail)
		v1.GET("/customers/by-username/:username", customerHandler.GetCustomerByUserName)
		
		// Student endpoints
		v1.POST("/students", studentHandler.Register)
		v1.GET("/students", studentHandler.GetAll)
		v1.GET("/students/by-email/:email", studentHandler.GetStudentByEmail)
		v1.GET("/students/by-username/:username", studentHandler.GetStudentByUserName)

		// auth endpoint
		v1.POST("/login", authHandler.GetToken)

	}

	return router
}

