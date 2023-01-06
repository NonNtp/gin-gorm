package v1

import (
	"github.com/NonNtp/gin-gorm/controllers"
	"github.com/NonNtp/gin-gorm/middlewares"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	//require jwt for all api
	// routerGroup := rg.Group("/users").Use(middlewares.AuthJWT())

	routerGroup := rg.Group("/users")

	routerGroup.GET("/", controllers.GetAll)
	routerGroup.GET("/search", controllers.SearchUsers)
	routerGroup.GET("/me", middlewares.AuthJWT(), controllers.GetProfile)
	routerGroup.GET("/:id", controllers.GetById)

	routerGroup.POST("/register", controllers.Register)
	routerGroup.POST("/login", controllers.Login)
}
