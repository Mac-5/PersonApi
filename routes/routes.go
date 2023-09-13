package routes

import (
	"github.com/Mac-5/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, personController *controllers.PersonController){
	personRoutes := router.Group("/api")
	
	personRoutes.GET("/", personController.GetAllPerson)
	personRoutes.GET("/:id", personController.GetOnePerson)
	personRoutes.POST("/", personController.CreatePerson)
	personRoutes.PUT("/:id", personController.UpdatePerson)
	personRoutes.DELETE("/:id", personController.DeletePerson)
}