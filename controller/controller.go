package controller

import (
	"MyProject/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
	router  *gin.Engine
}

func New(service *service.Service) *Controller {
	return &Controller{
		service: service,
		router:  gin.Default(),
	}

}
