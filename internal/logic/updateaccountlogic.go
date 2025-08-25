package logic

import (
	"context"

	"go-zero-api-demo/internal/svc"
	"go-zero-api-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountLogic {
	return &UpdateAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountLogic) UpdateAccount(req *types.UpdateAccountReq) (resp *types.Account, err error) {
	// todo: add your logic here and delete this line

	return
}
