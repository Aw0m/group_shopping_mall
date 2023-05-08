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

// user
type (
	// GetUserCountReq 获取用户数量
	GetUserCountReq struct {
	}

	GetUserListReq struct {
		PageNum  int `json:"page_num" validate:"required,gte=1"`
		PageSize int `json:"page_size" validate:"required,gte=1"`
	}
)

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

	GetCommodityCountReq struct {
	}

	GetCommodityFromIDReq struct {
		CommodityId int64 `json:"commodity_id,string"`
	}

	UpdateCommodityReq struct {
		CommodityId     int64    `json:"commodity_id,string" validate:"required,gte=1"`
		CommodityName   *string  `json:"commodity_name"`
		Price           *float64 `json:"price,string"`
		CategoryID      *int64   `json:"category_id,string"`
		CommodityStatus *int     `json:"commodity_status"`
		IsDeleted       *bool    `json:"is_deleted"`
	}
)

// category
type (
	// GetCategoryListReq 获取分类列表
	GetCategoryListReq struct {
	}

	// AddCategoryReq 添加分类
	AddCategoryReq struct {
		CategoryName string `json:"category_name" validate:"required,max=16"`
	}

	// DeleteCategoryReq 删除分类
	DeleteCategoryReq struct {
		CategoryId int64 `json:"category_id,string" validate:"required,gte=1"`
	}
)

// order
type (
	OrderCommodityInfo struct {
		CommodityId int64 `json:"commodity_id,string" validate:"required,gte=1"`
		Quantity    int   `json:"quantity" validate:"required,gte=1"`
	}
	// CreateOrderReq 创建订单
	CreateOrderReq struct {
		OrderCommodityInfoList []OrderCommodityInfo `json:"order_commodity_info_list" validate:"required,gte=1"` // 商品列表
	}

	ConfirmOrderReq struct {
		OrderId int64 `json:"order_id,string" validate:"required,gte=1"`
	}

	ShipOrderReq struct {
		OrderId             int64  `json:"order_id,string" validate:"required,gte=1"`
		DeliverymanName     string `json:"deliveryman_name" validate:"required,max=16"`
		DeliverymanPhoneNum string `json:"deliveryman_phone_num" validate:"required,max=16"`
	}

	// OrderReceiptReq 订单确认收获
	OrderReceiptReq struct {
		OrderId int64 `json:"order_id,string" validate:"required,gte=1"`
	}

	// GetOrderStatisticsReq 获取用户各种status订单的数量
	GetOrderStatisticsReq struct {
	}

	// GetOrderListReq 获取订单列表
	GetOrderListReq struct {
		PageSize   int   `json:"page_size"`
		PageNum    int   `json:"page_num"`
		StatusList []int `json:"status_list" validate:"required,gte=1"`
	}

	// GetOrderDetailReq 获取订单详情
	GetOrderDetailReq struct {
		OrderId int64 `json:"order_id,string" validate:"required,gte=1"`
	}

	// CancelOrderReq 取消订单
	CancelOrderReq struct {
		OrderId int64 `json:"order_id,string" validate:"required,gte=1"`
	}

	// GetAllOrderReq 获取所有订单列表
	GetAllOrderReq struct {
		PageSize   int   `json:"page_size"`
		PageNum    int   `json:"page_num"`
		StatusList []int `json:"status_list" validate:"required,gte=1"`
	}

	// GetOrderCountReq 获取订单数量
	GetOrderCountReq struct {
		StatusList []int `json:"status_list" validate:"required,gte=1"`
	}

	UpdateOrderInfoReq struct {
		OrderId             int64   `json:"order_id,string" validate:"required,gte=1"`
		SellerName          string  `json:"seller_name"`
		CommodityName       string  `json:"commodity_name"`
		DeliverymanName     string  `json:"deliveryman_name"`
		DeliverymanPhoneNum string  `json:"deliveryman_phone_num"`
		PhoneNum            string  `json:"phone_num"`
		ConsigneeName       string  `json:"consignee_name"`
		AddressDetailInfo   string  `json:"address_detail_info"`
		Price               float64 `json:"price" validate:"required,gte=1"`
		Quantity            int     `json:"quantity" validate:"required,gte=1"`
		Status              int     `json:"status" validate:"required,gte=1,lte=4"`
		IsDeleted           bool    `json:"is_deleted"`
	}
)
