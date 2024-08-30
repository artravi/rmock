package tests

import (
    "rmock/services/user"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
    router := gin.Default()
    router.GET("/users", user.GetUsers)

    req, _ := http.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "John Doe")
}

func TestGetUserByID(t *testing.T) {
    router := gin.Default()
    router.GET("/users/:id", user.GetUserByID)

    req, _ := http.NewRequest("GET", "/users/1", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "John Doe")
}

func TestGetUserByIDNotFound(t *testing.T) {
    router := gin.Default()
    router.GET("/users/:id", user.GetUserByID)

    req, _ := http.NewRequest("GET", "/users/999", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
    assert.Contains(t, w.Body.String(), "User not found")
}
