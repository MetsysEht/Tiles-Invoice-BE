package users

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Imanager interface {
	Create(*CreateRequest) error
	Login(*LoginRequest) (*LoginResponse, error)
	GetAll() (*GetUserArray, error)
	Update(*CreateRequest) error
	Delete(*DeleteRequest) error
}

type IRepo interface {
	Save(*User) error
	GetByUsername(string) (*User, error)
	GetAll() (*[]User, error)
	Update(*User) error
	Delete(string) error
}
