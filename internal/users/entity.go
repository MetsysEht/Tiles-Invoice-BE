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
	Role     string
}

type User struct {
	Id        string
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type GetUser struct {
	Id        string
	Username  string
	Role      string
	UpdatedAt time.Time
}

type GetUserArray []GetUser

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

func FromModelArray(uArray *[]models.User) *[]User {
	userArray := make([]User, len(*uArray))
	for u := range *uArray {
		userArray[u] = *FromModel(&(*uArray)[u])
	}
	return &userArray
}
