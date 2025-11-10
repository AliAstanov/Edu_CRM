package handler

import (
	"github.com/AliAstanov/Edu_CRM/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.ServiceI
}

func NewHandler(service service.ServiceI) handler {
	return handler{service: service}
}
func (h *handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Pong"})
}
