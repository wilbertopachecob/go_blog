package repository

import "wilbertopachecob/go_blog/cmd/web/models"

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
}
