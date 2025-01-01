package middleware

import (
	"strings"

	"github.com/Real-Musafir/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateRequest(bypass bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token != "" {
			token = strings.Split(token, " ")[1]
		}

		if token == "" || token == "null" || token == "undefined" {
			token = ctx.Request.Header.Get("token")
		}

		if token == "" || token == "null" || token == "undefined" {
			token, _ = ctx.Cookie("token")
		}

		if token == "" {
			token, _ = ctx.GetQuery("token")
		}

		if bypass && token == "" {
			ctx.Next()
			return
		}

		if token == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized, Please login again",
				"statusCode": 401,
			})

			ctx.Abort()
			
		}else {
			jwtToken, err := utils.VerifyToken(token)

				if err != nil {
						ctx.JSON(401, gin.H{
						"message": "Unauthorized, Please login again",
						"statusCode": 401,
					})
					ctx.Abort()				
				}else {
					if user_id, ok := (*jwtToken)["user_id"]; ok {
						
						user_name := (*jwtToken)["user_name"]

						ctx.Set("user_id", user_id)
						ctx.Set("user_name", user_name)

						ctx.Next()
					} 
				}
		}
	}
}