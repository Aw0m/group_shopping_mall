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
	OrderStatus_Pending   = 1 // 待确认（未付款）
	OrderStatus_Unshipped = 2 // 待发货（已付款）
	OrderStatus_Shipped   = 3 // 已发货
)
