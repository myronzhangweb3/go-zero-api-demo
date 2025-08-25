package logic

import (
	"context"

	"go-zero-api-demo/internal/svc"
	"go-zero-api-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountLogic) GetAccount(req *types.IdReq) (resp *types.Account, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.AccountModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp = new(types.Account)
	resp.Id = res.Id
	resp.Address = res.Address
	resp.AddressIndex = res.AddressIndex
	resp.Comment = res.Comment.String
	resp.CreatedAt = res.CreatedAt.Time.String()
	resp.UpdatedAt = res.UpdatedAt.Time.String()
	resp.DeletedAt = res.DeletedAt.Time.String()

	return resp, nil
}
