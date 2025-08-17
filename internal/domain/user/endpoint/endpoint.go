package endpoint

import (
	"architecture/internal/domain/model"
	user_dto "architecture/internal/domain/user/dto"
	user_error "architecture/internal/domain/user/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Service interface {
		FindUserByID(id string) (*model.User, error)
		GetUsers() *[]model.User
		Create(user user_dto.CreateUserStruct) (*model.User, error)
	}

	UserEndpoint struct {
		service Service
	}
)

func New(service Service) *UserEndpoint {
	return &UserEndpoint{
		service: service,
	}
}

func (ue UserEndpoint) Route(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.GET("/:id", ue.GetUser)
	userGroup.GET("/", ue.GetUsers)
	userGroup.POST("/create", ue.CreateUser)
}

func (ue *UserEndpoint) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := ue.service.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": user_error.ErrUserNotFound})
		return
	}

	c.JSON(200, user)
}

func (ue *UserEndpoint) GetUsers(c *gin.Context) {
	users := ue.service.GetUsers()

	c.JSON(http.StatusOK, users)
}

func (ue *UserEndpoint) CreateUser(c *gin.Context) {
	var user user_dto.CreateUserStruct

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ue.service.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
