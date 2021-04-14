package crud

import (
	"wilbertopachecob/go_blog/cmd/web/channels"
	"wilbertopachecob/go_blog/cmd/web/models"

	"gorm.io/gorm"
)

type repositoryUserCRUD struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryUserCRUD {
	return &repositoryUserCRUD{db}
}

func (r *repositoryUserCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, nil
}

func (r *repositoryUserCRUD) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return nil, nil
}
