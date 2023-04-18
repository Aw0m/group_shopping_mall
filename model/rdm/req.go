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

type GetCartListReq struct {
}
