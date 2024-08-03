package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/normatov07/mini-tweet/app"
	"github.com/normatov07/mini-tweet/common/response"
	"github.com/normatov07/mini-tweet/common/token"
)

func AuthMiddleware(app *app.ApplicationContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenMaker, err := token.NewPasetoMaker()
		if err != nil {
			log.Panic(err)
		}
		authorizationHeader := ctx.GetHeader("authorization")
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse("authorization header is not provided", 403))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse("invalid authorization header format", 403))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse("unsupported authorization type", 403))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err.Error(), 403))
			return
		}

		app.User = payload
		ctx.Next()
	}
}
