package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/NonNtp/gin-gorm/db"
	"github.com/NonNtp/gin-gorm/dto"
	"github.com/NonNtp/gin-gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthJWT() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenHeader := ctx.GetHeader("Authorization")
		accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]
		// fmt.Println(accessToken)
		jwtSecretKey := os.Getenv("JWT_SECRET")

		token, _ := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			claims := token.Claims.(jwt.MapClaims)
			var user models.User
			db.Conn.First(&user, claims["user_id"])
			ctx.Set("user", dto.UserResponse{
				ID:       user.ID,
				Fullname: user.Fullname,
				Email:    user.Email,
			})
			// return to next method if token is exist
			ctx.Next()
		}

	})
}
