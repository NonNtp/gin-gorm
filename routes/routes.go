package routes

import (
	v1 "github.com/NonNtp/gin-gorm/routes/v1"
	v2 "github.com/NonNtp/gin-gorm/routes/v2"

	"github.com/gin-gonic/gin"
)

func ServeRoutes(r *gin.Engine) {
	apiV1 := r.Group("/api/v1")
	apiV2 := r.Group("/api/v2")

	v1.InitUserRoutes(apiV1)

	v2.InitHomeRoutes(apiV2)

}
