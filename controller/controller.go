package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prankevich/MyProject/service"
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
