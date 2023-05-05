package bdm

import "time"

type Order struct {
	OrderId             int64     `json:"order_id,string"`
	SellerId            int64     `json:"seller_id,string"`
	SellerName          string    `json:"seller_name"`
	CommodityId         int64     `json:"commodity_id,string"`
	CommodityName       string    `json:"commodity_name"`
	CommodityImageURL   string    `json:"commodity_image_url"`
	CustomerId          int64     `json:"customer_id,string"`
	DeliverymanName     string    `json:"deliveryman_name"`
	DeliverymanPhoneNum string    `json:"deliveryman_phone_num"`
	PhoneNum            string    `json:"phone_num"`
	ConsigneeName       string    `json:"consignee_name"`
	AddressId           int64     `json:"address_id"`
	AddressDetailInfo   string    `json:"address_detail_info"`
	Price               float64   `json:"price"`
	Quantity            int       `json:"quantity"`
	Status              int       `json:"status"`
	CreateTime          time.Time `json:"create_time"`
	UpdateTime          time.Time `json:"update_time"`
	IsDeleted           bool      `json:"is_deleted"`
}
