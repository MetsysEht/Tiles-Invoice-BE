package users

import (
	"time"

	"github.com/MetsysEht/Tiles-Invoice-BE/internal/database/models"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/uniqueId"
	"github.com/MetsysEht/Tiles-Invoice-BE/utils"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JwtToken string
}

type User struct {
	Id        string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ToModel() *models.User {
	if utils.IsEmpty(u.Id) {
		u.Id = uniqueId.New()
	}
	return &models.User{
		ID:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}

func FromModel(u *models.User) *User {
	user := &User{
		Id:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	return user
}
