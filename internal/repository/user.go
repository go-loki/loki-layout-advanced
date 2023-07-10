package repository

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx *app.RequestContext, user *model.User) error
	Update(ctx *app.RequestContext, user *model.User) error
	GetByID(ctx *app.RequestContext, id string) (*model.User, error)
	GetByUsername(ctx *app.RequestContext, username string) (*model.User, error)
}

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}
func (r *userRepository) Create(ctx *app.RequestContext, user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *userRepository) Update(ctx *app.RequestContext, user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (r *userRepository) GetByID(ctx *app.RequestContext, userId string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {

		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return &user, nil
}

func (r *userRepository) GetByUsername(ctx *app.RequestContext, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return &user, nil
}
