package rdm

type WXLoginRequest struct {
	Code      string `json:"code"` // 临时登录凭证。通过code在微信服务器获取openid
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Gender    string `json:"gender"` // m/f/u
	PhoneNum  string `json:"phone_num"`
}

type GetUserinfoReq struct {
}

type UpdateUserinfoReq struct {
	Username  string `json:"username" validate:"max=16"`
	Gender    string `json:"gender"`
	AvatarURL string `json:"avatar_url"`
	PhoneNum  string `json:"phone_num"`
	AddressId int64  `json:"address_id,string"`
}

type GetAddressListReq struct {
}

type AddShoppingAddressReq struct {
	AddressName string `json:"address_name"`
	DetailInfo  string `json:"detail_info"`
}

// cart
type (
	// GetCartListReq 获取购物车列表
	GetCartListReq struct {
	}

	UpdateCartNumReq struct {
		CartId   int64 `json:"cart_id,string" validate:"required,gte=1"`
		Quantity int   `json:"quantity" validate:"required,gte=1"`
	}

	DeleteCartReq struct {
		CartId int64 `json:"cart_id,string" validate:"required,gte=1"`
	}
)

// commodity
type (
	// GetCommodityDetailReq 获取商品详情
	GetCommodityDetailReq struct {
		CommodityId int64 `json:"commodity_id,string" validate:"required,gte=1"`
	}

	AddToCartReq struct {
		CommodityId  int64 `json:"commodity_id,string" validate:"required,gte=1"`
		CommodityNum int   `json:"commodity_num" validate:"required,gte=1"`
	}

	GetCommodityFromCategoryReq struct {
		CategoryId int64 `json:"category_id,string"`
		PageSize   int   `json:"page_size"`
		PageNum    int   `json:"page_num"`
	}
)

// category
type (
	// GetCategoryListReq 获取分类列表
	GetCategoryListReq struct {
	}
)
