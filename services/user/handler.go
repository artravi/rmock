package user

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

var users = []User{
    {ID: 1, Name: "John Doe"},
    {ID: 2, Name: "Jane Doe"},
}

func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
