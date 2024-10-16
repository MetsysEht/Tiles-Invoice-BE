package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	manager Imanager
}

func CreateServer(manager Imanager) IServer {
	return Server{manager}
}

func (s Server) Create(c *gin.Context) {
	req := CreateRequest{}
	err := c.Bind(&req)
	if err != nil {
		return
	}
	s.manager.Create(&req)
	c.JSON(200, gin.H{})
}

func (s Server) Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return
	}
	resp, err := s.manager.Login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("jwt", resp.JwtToken, int(24*time.Second.Seconds()), "/", "localhost", true, true)
	c.JSON(200, resp)
	return
}

func (s Server) GetAll(c *gin.Context) {
	resp, err := s.manager.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}
