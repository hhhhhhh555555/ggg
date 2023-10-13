package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/app/goods/internal/biz"
)

type GoodsDataCase struct {
	data *Data
	log  *log.Helper
}

func NewGoodsDataCase(data *Data, logger log.Logger) biz.GoodsMethods {
	return &GoodsDataCase{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (g *GoodsDataCase) Add(ctx context.Context, b *biz.Goods) (*biz.Goods, error) {
	return b, nil
}
