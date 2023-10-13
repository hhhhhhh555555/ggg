package service

import (
	"context"
	v1 "helloworld/api/goods/v1"
	"helloworld/app/goods/internal/biz"
)

type GoodsMethods struct {
	v1.UnimplementedGoodsServer
	uc *biz.GoodsMethodCase
}

func NewGoodService(uc *biz.GoodsMethodCase) *GoodsMethods {
	return &GoodsMethods{uc: uc}
}
func (s *GoodsMethods) Add(ctx context.Context, in *v1.AddGoodsRequest) (*v1.AddGoodsResponse, error) {
	_, err := s.uc.CreateGoods(ctx, &biz.Goods{
		Id:          in.Id,
		Name:        in.Name,
		CategoryId:  in.CategoryId,
		Price:       in.Price,
		Image:       in.Image,
		Description: in.Description,
		Status:      in.Status,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		CreateUser:  in.CreateUser,
		UpdateUser:  in.UpdateUser,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddGoodsResponse{
		Code:    200,
		Success: "添加成功",
	}, nil
}
