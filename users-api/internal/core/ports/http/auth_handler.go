package httpport

import "github.com/gin-gonic/gin"


type AuthHandler interface{
	GetToken(ctx *gin.Context) 
}