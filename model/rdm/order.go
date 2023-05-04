package rdm

import (
	"time"
)

// Order 订单信息表
type Order struct {
	OrderId             int64     `gorm:"column:order_id;type:bigint(20);AUTO_INCREMENT;primary_key" json:"order_id,string"`
	SellerId            int64     `gorm:"column:seller_id;type:bigint(20);NOT NULL" json:"seller_id,string"`
	SellerName          string    `gorm:"column:seller_name;type:varchar(16);NOT NULL" json:"seller_name"`
	CommodityId         int64     `gorm:"column:commodity_id;type:bigint(20);NOT NULL" json:"commodity_id,string"`
	CommodityName       string    `gorm:"column:commodity_name;type:varchar(16);NOT NULL" json:"commodity_name"`
	CustomerId          int64     `gorm:"column:customer_id;type:bigint(20);NOT NULL" json:"customer_id,string"`
	DeliverymanName     string    `gorm:"column:deliveryman_name;type:varchar(16)" json:"deliveryman_name"`
	DeliverymanPhoneNum string    `gorm:"column:deliveryman_phone_num;type:varchar(16);comment:送货人的电话号码" json:"deliveryman_phone_num"`
	PhoneNum            string    `gorm:"column:phone_num;type:varchar(16);comment:收货人的电话号码;NOT NULL" json:"phone_num"`
	ConsigneeName       string    `gorm:"column:consignee_name;type:varchar(16);comment:收货人的收货名;NOT NULL" json:"consignee_name"`
	AddressId           int64     `gorm:"column:address_id;type:bigint(20);NOT NULL" json:"address_id"`
	AddressDetailInfo   string    `gorm:"column:address_detail_info;type:varchar(255)" json:"address_detail_info"`
	Price               float64   `gorm:"column:price;type:double;NOT NULL" json:"price"`
	Quantity            int       `gorm:"column:quantity;type:int(11);NOT NULL" json:"quantity"`
	Status              int       `gorm:"column:status;type:int(11);default:1;NOT NULL" json:"status"`
	CreateTime          time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime          time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	IsDeleted           bool      `gorm:"column:is_deleted;type:tinyint(1);default:0;comment:是否删除;NOT NULL" json:"is_deleted"`
}

func (m *Order) TableName() string {
	return "order"
}
