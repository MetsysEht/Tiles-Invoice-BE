package users

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = s.manager.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func (s Server) Login(c *gin.Context) {
	req := LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := s.manager.Login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

func (s Server) GetAll(c *gin.Context) {
	resp, err := s.manager.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

func (s Server) Update(c *gin.Context) {
	req := CreateRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = s.manager.Update(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func (s Server) Delete(c *gin.Context) {
	req := DeleteRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = s.manager.Delete(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}
