package controllers

import (
	"net/http"

	"github.com/NonNtp/gin-gorm/db"
	"github.com/NonNtp/gin-gorm/dto"
	"github.com/NonNtp/gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func CreateBlog(ctx *gin.Context) {
	var form dto.BlogRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := models.Blog{
		Topic: form.Topic,
		Detail: form.Detail,
	}

	if err := db.Conn.Create(&blog).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated , dto.BlogResponse{
		ID: blog.ID,
		Topic: blog.Topic,
		Detail: blog.Detail,
	})

}
