package bdm

import "time"

type Cart struct {
	CartId      int64     `json:"cart_id,string"`
	CustomerId  int64     `json:"customer_id,string"`
	CommodityId int64     `json:"commodity_id,string"`
	InitPrice   float64   `json:"init_price"`
	CreateTime  time.Time `json:"create_time"`
	IsDeleted   bool      `json:"is_deleted"`
}
