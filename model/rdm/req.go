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
