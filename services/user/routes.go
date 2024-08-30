package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", GetUsers)
        userRoutes.GET("/:id", GetUserByID)
    }
}
