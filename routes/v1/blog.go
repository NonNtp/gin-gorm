package v1

import (
	"github.com/NonNtp/gin-gorm/controllers"
	"github.com/NonNtp/gin-gorm/middlewares"
	"github.com/gin-gonic/gin"
)

func InitBlogRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/blog")

	routerGroup.POST("/:id", middlewares.AuthJWT(), controllers.CreateBlog)
}
