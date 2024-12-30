package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0]
			errMap := make(map[string]any, 0)

			if err != nil {
				splittedError := strings.Split(err.Error(), "::")
				fmt.Printf("Method Name : %s, Status Code: %s, Client Message : %s, Error : %v", splittedError[2], splittedError[0], splittedError[1], splittedError[3])

				statusCode, err := strconv.Atoi(splittedError[0])
				if err != nil {
					statusCode = 500
				}

				errMap["message"] = splittedError[1]
				errMap["statusCode"] = statusCode
				ctx.JSON(statusCode, errMap)

			}

			
		}
	}
}