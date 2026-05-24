package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/medurga/practice-go/src/handler"
	"github.com/medurga/practice-go/src/print"
)

func main() {
	fmt.Println("Hello, World!")
	print.PrintMessage()

	r := gin.Default()
	r.GET("/ping", handler.PongHandler)
	r.GET("/user", handler.GetUser)

	r.Run()
}
