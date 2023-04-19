package bdm

import "time"

type Seller struct {
	SellerId   int64     `json:"seller_id,string"`
	SellerName string    `json:"seller_name"`
	PhoneNum   string    `json:"phone_num"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	IsDeleted  bool      `json:"is_deleted"`
}
