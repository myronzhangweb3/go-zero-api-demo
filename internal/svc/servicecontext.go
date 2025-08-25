package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-api-demo/internal/config"
	"go-zero-api-demo/internal/model"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model.AccountModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:       c,
		AccountModel: model.NewAccountModel(conn),
	}
}
