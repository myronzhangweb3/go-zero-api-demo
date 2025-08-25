package logic

import (
	"context"

	"go-zero-api-demo/internal/svc"
	"go-zero-api-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAccountsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAccountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAccountsLogic {
	return &ListAccountsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAccountsLogic) ListAccounts() (resp []types.Account, err error) {
	// todo: add your logic here and delete this line

	return
}
