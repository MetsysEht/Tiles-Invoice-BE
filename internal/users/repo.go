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

func (r Repo) GetAll() (*[]User, error) {
	var users []models.User
	q := r.DB.Find(&users)
	err := gormDatabase.GetDatabaseError(q)
	if err != nil {
		return nil, err
	}
	return FromModelArray(&users), nil
}

func (r Repo) Save(user *User) error {
	userModel := user.ToModel()
	q := r.DB.Save(userModel)
	err := gormDatabase.GetDatabaseError(q)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) Update(user *User) error {
	userModel := user.ToModel()
	q := r.DB.Model(userModel).Updates(userModel)
	err := gormDatabase.GetDatabaseError(q)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) Delete(username string) error {
	q := r.DB.Delete(models.User{}, "username = ?", username)
	err := gormDatabase.GetDatabaseError(q)
	if err != nil {
		return err
	}
	return nil
}
