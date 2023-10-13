package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"helloworld/app/goods/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysql, NewGoodsDataCase)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
func NewMysql(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	if err := db.AutoMigrate(&Goods{}); err != nil {
		log.Fatal(err)
	}
	return db
}

type Goods struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'主键'"`
	Name        string    `gorm:"column:name;NOT NULL;comment:'菜品名称'"`
	CategoryId  int64     `gorm:"column:category_id;NOT NULL;comment:'菜品分类id'"`
	Price       string    `gorm:"column:price;default:NULL;comment:'菜品价格'"`
	Image       string    `gorm:"column:image;default:NULL;comment:'图片'"`
	Description string    `gorm:"column:description;default:NULL;comment:'描述信息'"`
	Status      int32     `gorm:"column:status;default:1;comment:'0 停售 1 起售'"`
	CreateTime  time.Time `gorm:"column:create_time;default:NULL;comment:'创建时间'"`
	UpdateTime  time.Time `gorm:"column:update_time;default:NULL;comment:'更新时间'"`
	CreateUser  int64     `gorm:"column:create_user;default:NULL;comment:'创建人'"`
	UpdateUser  int64     `gorm:"column:update_user;default:NULL;comment:'修改人'"`
}

func (d *Goods) TableName() string {
	return "goods"
}
