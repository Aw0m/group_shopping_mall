package rdm

import "group_shopping_mall/model/bdm"

type WXLoginResponse struct {
	Token string   `json:"token"`
	User  bdm.User `json:"user"`
}

type GetUserinfoRsp struct {
	User bdm.User `json:"user"`
}

type UpdateUserinfoRsp struct {
}

type GetAddressListRsp struct {
	AddressList []bdm.ShoppingAddress `json:"address_list"`
}

type AddShoppingAddressRsp struct {
}

type (
	CartCommodityItem struct {
		Commodity bdm.Commodity `json:"commodity"`
		Cart      bdm.Cart      `json:"cart"`
	}
	CartInfoItem struct {
		SellerID      int64               `json:"seller_id,string"`
		SellerName    string              `json:"seller_name"`
		CommodityList []CartCommodityItem `json:"commodity_list"`
	}
	GetCartListRsp struct {
		CommodityItemList []CartInfoItem `json:"commodity_item_list"`
	}
)
