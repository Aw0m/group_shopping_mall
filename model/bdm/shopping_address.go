package bdm

// ShoppingAddress 自提点表
type ShoppingAddress struct {
	AddressId   int64  `json:"address_id,string"`
	AddressName string `json:"address_name"`
	DetailInfo  string `json:"detail_info"`
}
