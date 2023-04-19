package rdm

import (
	"time"
)

// Cart 购物车表
type Cart struct {
	CartId       int64     `gorm:"column:cart_id;type:bigint(20);AUTO_INCREMENT;comment:购物车的id;primary_key" json:"cart_id"`
	CustomerId   int64     `gorm:"column:customer_id;type:bigint(20);comment:用户id;NOT NULL" json:"customer_id"`
	CommodityId  int64     `gorm:"column:commodity_id;type:bigint(20);comment:商品分类的名字;NOT NULL" json:"commodity_id"`
	InitPrice    float64   `gorm:"column:init_price;type:double;comment:初始加入购物车时商品的价格;NOT NULL" json:"init_price"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除" json:"is_deleted"`
	CommodityNum int       `gorm:"column:commodity_num;type:int(11);comment:加入购物车的数量;NOT NULL" json:"commodity_num"`
}

func (m *Cart) TableName() string {
	return "cart"
}
