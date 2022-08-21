package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-short/internal/config"
	"go-zero-short/internal/model/short"
	"go-zero-short/internal/model/users"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel users.UsersModel
	ShortModel short.ShortModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	usersModel := users.NewUsersModel(coon, c.CacheRedis)
	shortModel := short.NewShortModel(coon, c.CacheRedis)
	return &ServiceContext{
		Config:     c,
		UsersModel: usersModel,
		ShortModel: shortModel,
	}
}
