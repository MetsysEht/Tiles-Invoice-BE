package users

import (
	"time"

	"github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	repo IRepo
}

func NewManager(repo IRepo) Imanager {
	return Manager{repo: repo}
}

func (m Manager) Create(req *CreateRequest) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &User{
		Username: req.Username,
		Password: string(passwordHash),
		Role:     req.Role,
	}
	m.repo.Save(user)
	return
}

func (m Manager) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := m.repo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"iss": "BE",                                  // Issuer
		"aud": user.Role,                             // Audience (user role)
		"exp": time.Now().Add(24 * time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                     // Issued at
	})
	token, err := claims.SignedString([]byte(boot.Config.App.Key))
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		JwtToken: token,
	}, nil
}
