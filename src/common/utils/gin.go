package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetMode() {
	mode := os.Getenv("APP_MODE")
	gin.SetMode(mode)
}
