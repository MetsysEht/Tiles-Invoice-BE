package users

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/database/models"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/gormDatabase"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r Repo) GetByUsername(username string) (*User, error) {
	user := models.User{}
	q := r.DB.Where("username = ?", username).First(&user)
	err := gormDatabase.GetDatabaseError(q)
	if err != nil {
		return nil, err
	}
	return FromModel(&user), nil
}

func (r Repo) GetAll() {
	//TODO implement me
	panic("implement me")
}

func (r Repo) Save(user *User) {
	userModel := user.ToModel()
	_ = r.DB.Save(userModel)
	return
}
