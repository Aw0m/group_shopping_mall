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

type (
	CommodityInfo struct {
		CommodityId   int64   `json:"commodity_id,string"`
		CommodityName string  `json:"commodity_name"`
		Price         float64 `json:"price"`
		ImageURL      string  `json:"image_url"`
	}
	SellerInfo struct {
		SellerId   int64  `json:"seller_id,string"`
		SellerName string `json:"seller_name"`
	}
	CommentInfo struct {
		CommentId          int64  `json:"comment_id,string"`    // 评论id
		CommenterName      string `json:"commenter_name"`       // 评论者名字
		CommenterAvatarURL string `json:"commenter_avatar_url"` // 评论者头像
		CommentText        string `json:"comment_text"`         // 评论内容
		ScoreRes           int    `json:"score_res"`            // 1-5星打分
	}
	GetCommodityDetailRsp struct {
		CommodityInfo CommodityInfo `json:"commodity_info"`
		SellerInfo    SellerInfo    `json:"seller_info"`
		CommentInfo   CommentInfo   `json:"comment_info"`
	}
)

type AddToCartRsp struct {
}

type UpdateCartNumRsp struct {
}

type GetCategoryListRsp struct {
	CategoryList []bdm.Category `json:"category_list"`
}

type GetCommodityFromCategoryRsp struct {
	CommodityList []bdm.Commodity `json:"commodity_list"`
}
