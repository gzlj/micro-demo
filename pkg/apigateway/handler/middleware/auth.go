package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("X-Token")
		if strings.Compare(token, "iloveyou") != 0 {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"message": "token is not valid",
				"code": 403,
			})
			return
		}
		ctx.Next()
	}
}

