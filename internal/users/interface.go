package users

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	GetAll(c *gin.Context)
}

type Imanager interface {
	Create(*CreateRequest)
	Login(*LoginRequest) (*LoginResponse, error)
	GetAll() (*GetUserArray, error)
}

type IRepo interface {
	Save(*User)
	GetByUsername(string) (*User, error)
	GetAll() (*[]User, error)
}
