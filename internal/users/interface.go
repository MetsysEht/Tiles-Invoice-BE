package users

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
}

type Imanager interface {
	Create(*CreateRequest)
	Login(request *LoginRequest) (*LoginResponse, error)
}

type IRepo interface {
	Save(*User)
	GetByUsername(string) (*User, error)
	GetAll()
}
