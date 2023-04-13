package constant

import (
	"fmt"

	"group_shopping_mall/utils/utils"
)

var (
	OpenIDURL string
)

func InitConstant(path string) {
	type GlobalConfig struct {
		AppId  string `json:"appid"`
		Secret string `json:"secret"`
	}
	config := utils.GetConfig[GlobalConfig](path)
	OpenIDURL = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%%s&grant_type=authorization_code",
		config.AppId,
		config.Secret,
	)
}

var MySigningKey = []byte("dev123")

const (
	Duration = 24 * 30
)

const (
	AdminType_NotAdmin   = 1
	AdminType_Admin      = 2
	AdminType_SuperAdmin = 3
)

const (
	OrderStatus_Unconfirmed = 1
	OrderStatus_UnReviewed  = 2
	OrderStatus_Reviewed    = 3
	OrderStatus_Deleted     = 4

	OrderReviewResult_Pass   = 1
	OrderReviewResult_Delete = 2
)

const (
	PriceType_Fixed = 1 // 1固定定价
	PriceType_Incr  = 2 // 2增量定价
	PriceType_Rate  = 3 // 3倍率定价
)

const (
	OrderFrom_Customer = 1 // 1 客户的订单
	OrderFrom_Playmate = 2 // 2 陪玩的订单
	OrderFrom_Club     = 3 // 3 俱乐部的订单
)

const (
	UserType_Technology    = 2 // 2 技术
	UserType_Entertainment = 1 // 1 娱乐
)

const (
	CommodityType_Gift  = 1 // 1 礼物
	CommodityType_Other = 2 // 2 其他

	OrderType_Gift  = 1 // 1 礼物
	OrderType_Other = 2 // 2 其他
)
