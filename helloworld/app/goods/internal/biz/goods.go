package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "helloworld/api/goods/v1"
)

var (
	ErrorGoodAddFailed = errors.NotFound(v1.ErrorReason_ADD_FAILED.String(), "goods add failed")
)

type Goods struct {
	Id          int64  `json:"id,omitempty"`          // 主键
	Name        string `json:"name,omitempty"`        // 菜品名称
	CategoryId  int64  `json:"category_id,omitempty"` // 菜品分类id
	Price       string `json:"price,omitempty"`       // 菜品价格
	Image       string `json:"image,omitempty"`       // 图片
	Description string `json:"description,omitempty"` // 描述信息
	Status      int32  `json:"status,omitempty"`      // 0 停售 1 起售
	CreateTime  string `json:"create_time,omitempty"` // 创建时间
	UpdateTime  string `json:"update_time,omitempty"` // 更新时间
	CreateUser  int64  `json:"create_user,omitempty"` // 创建人
	UpdateUser  int64  `json:"update_user,omitempty"` // 修改人
}
type GoodsMethods interface {
	Add(context.Context, *Goods) (*Goods, error)
}
type GoodsMethodCase struct {
	method GoodsMethods
	log    *log.Helper
}

func NewGoodsMethodCase(method GoodsMethods, logger log.Logger) *GoodsMethodCase {
	return &GoodsMethodCase{
		method: method,
		log:    log.NewHelper(logger),
	}
}
func (this *GoodsMethodCase) CreateGoods(ctx context.Context, g *Goods) (*Goods, error) {
	this.log.WithContext(ctx).Infof("添加一件商品哦")
	return this.method.Add(ctx, g)
}
