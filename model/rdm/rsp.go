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
