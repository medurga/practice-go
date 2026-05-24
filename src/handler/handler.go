package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/medurga/practice-go/src/model"
)

func PongHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.User{
		Id:   100,
		Name: "Durga",
	})
}
