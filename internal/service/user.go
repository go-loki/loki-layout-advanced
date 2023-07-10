package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/model"
	"github.com/go-hasaki/hasaki-layout-advanced/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required,email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type UserService interface {
	Register(ctx *app.RequestContext, req *RegisterRequest) error
	Login(ctx *app.RequestContext, req *LoginRequest) (string, error)
	GetProfile(ctx *app.RequestContext, userId string) (*model.User, error)
	UpdateProfile(ctx *app.RequestContext, userId string, req *UpdateProfileRequest) error
	GenerateToken(ctx *app.RequestContext, userId string) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

func (s *userService) Register(ctx *app.RequestContext, req *RegisterRequest) error {
	// 检查用户名是否已存在
	if user, err := s.userRepo.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	// 生成用户ID
	userId, err := s.sid.GenString()
	if err != nil {
		return errors.Wrap(err, "failed to generate user ID")
	}
	// 创建用户
	user := &model.User{
		UserId:   userId,
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return nil
}

func (s *userService) Login(ctx *app.RequestContext, req *LoginRequest) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", errors.Wrap(err, "failed to get user by username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	// 生成JWT token
	token, err := s.GenerateToken(ctx, user.UserId)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}

func (s *userService) GetProfile(ctx *app.RequestContext, userId string) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}

	return user, nil
}

func (s *userService) UpdateProfile(ctx *app.RequestContext, userId string, req *UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return errors.Wrap(err, "failed to get user by ID")
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (s *userService) GenerateToken(ctx *app.RequestContext, userId string) (string, error) {
	// 生成JWT token
	token, err := s.jwt.GenToken(userId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate JWT token")
	}

	return token, nil
}