package delivery

import (
	"net/http"
	"project/internal/entity"
	UseCase "project/internal/usecase"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userUsecase UseCase.UserUsecase
}

// NewUserHandler initializes routes for user operations
func NewUserHandler(router *gin.Engine, userUsecase UseCase.UserUsecase) {
	handler := &UserHandler{userUsecase: userUsecase}

	router.POST("/users", handler.CreateUser) // Create user route
	router.GET("/users", handler.GetAllUsers) // Get all users route
}

// CreateUser handles HTTP POST for creating a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userUsecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetAllUsers handles HTTP GET for retrieving all users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
