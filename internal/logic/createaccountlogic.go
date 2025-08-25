package logic

import (
	"context"

	"go-zero-api-demo/internal/svc"
	"go-zero-api-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountLogic {
	return &CreateAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAccountLogic) CreateAccount(req *types.CreateAccountReq) (resp *types.Account, err error) {
	// todo: add your logic here and delete this line

	return
}
