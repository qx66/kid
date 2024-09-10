package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
		return
	}
	
	g := gin.New()
	
	g.POST("/")
	err = g.Run(":8080")
	
	if err != nil {
		logger.Error("程序异常", zap.Error(err))
	}
}
