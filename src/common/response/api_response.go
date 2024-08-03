package response

import "github.com/gin-gonic/gin"

func SuccessResponse(data any) gin.H {
	return gin.H{
		"success": true,
		"result":  data,
	}
}

func ErrorResponse(err string, code int) gin.H {
	return gin.H{
		"success": false,
		"error": gin.H{
			"message": err,
			"code":    code,
		},
	}
}
