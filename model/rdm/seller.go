package rdm

import "time"

// Seller 购物车表
type Seller struct {
	SellerId   int64     `gorm:"column:seller_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"seller_id"`
	SellerName string    `gorm:"column:seller_name;type:varchar(16);NOT NULL" json:"seller_name"`
	PhoneNum   string    `gorm:"column:phone_num;type:varchar(16);NOT NULL" json:"phone_num"`
	Status     int       `gorm:"column:status;type:int(11);default:1;NOT NULL" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	IsDeleted  bool      `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除;NOT NULL" json:"is_deleted"`
}

func (m *Seller) TableName() string {
	return "seller"
}
