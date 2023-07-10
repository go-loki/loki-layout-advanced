package service

import (
	"github.com/go-hasaki/hasaki-layout-advanced/internal/middleware"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/helper/sid"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *middleware.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
