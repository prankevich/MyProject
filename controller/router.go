package controller

import "github.com/gin-gonic/gin"

func (ctrl *Controller) RegisterEndpoints() {
	ctrl.router.GET("/ping", ctrl.Ping)

	ctrl.router.GET("/users", ctrl.GetAllUsers)
	ctrl.router.GET("/users/:id", ctrl.GetUsersByID)
	ctrl.router.POST("/users", ctrl.CreateUsersByID)
	ctrl.router.PUT("/users/:id", ctrl.UpdateUserByID)
	ctrl.router.DELETE("/users/:id", ctrl.DeleteUserByID)
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}
func (ctrl *Controller) RunServer(address string) error {
	ctrl.RegisterEndpoints()
	if err := ctrl.router.Run(address); err != nil {
		return err
	}
	return nil

}
