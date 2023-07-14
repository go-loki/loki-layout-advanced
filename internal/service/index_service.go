package service

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type IndexService interface {
	IndexV1() map[string]any
}

func NewIndexService() IndexService {
	return &indexService{}
}

type indexService struct{}

func (svc *indexService) IndexV1() map[string]any {
	name := ""
	age := 0
	userInfo := viper.GetStringMap("user_info")
	if f, ok := userInfo["name"]; ok && f != "" {
		name = cast.ToString(f)
	}
	if f, ok := userInfo["age"]; ok && f != "" {
		age = cast.ToInt(f)
	}
	return map[string]any{
		"name": name,
		"age":  age,
	}
}
