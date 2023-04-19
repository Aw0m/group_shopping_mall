package rdm

import "gorm.io/datatypes"

// Commodity 商品信息表
type Commodity struct {
	CommodityId     int64          `gorm:"column:commodity_id;type:bigint(20);AUTO_INCREMENT;comment:商品的id;primary_key" json:"commodity_id"`
	CommodityName   string         `gorm:"column:commodity_name;type:varchar(16);comment:商品的名字" json:"commodity_name"`
	Price           float64        `gorm:"column:price;type:double;default:9999;comment:商品的价格" json:"price"`
	SellerId        int64          `gorm:"column:seller_id;type:bigint(20);comment:卖家的id" json:"seller_id"`
	CategoryId      int64          `gorm:"column:category_id;type:bigint(20);comment:商品所属分类;NOT NULL" json:"category_id"`
	Content         datatypes.JSON `gorm:"column:content;type:json;comment:商品的详细内容" json:"content"`
	AddressList     string         `gorm:"column:address_list;type:varchar(255);comment:支持配送地址" json:"address_list"`
	CommodityStatus int            `gorm:"column:commodity_status;type:int(11);default:1;comment:商品状态" json:"commodity_status"`
	IsDeleted       bool           `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除" json:"is_deleted"`
}

func (m *Commodity) TableName() string {
	return "commodity"
}
