package biz

import (
	"context"
	"fmt"
	"net/http"

	"group_shopping_mall/dal/dao"

	"github.com/pkg/errors"

	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/utils"
)

// WXLogin 用户登录服务，生成Token。并且在数据库中检索该用户是否已经注册，如果没有则还会在数据库中创建该用户
func WXLogin(ctx context.Context, code, username, avatarURL string, gender string) (u bdm.User, token string, retErr error) {
	url := fmt.Sprintf(constant.OpenIDURL, code)
	// 通过code获取用户的唯一标识符openid
	res, err := http.Get(url)
	if err != nil {
		return u, "", errors.Wrap(err, "fail to get openid")
	}

	// 解析微信服务端的response，获得openid并查询是否已经存入数据库，如果没有则在数据库中生成一个user
	body, _ := utils.ParseResponse(res)
	openidAny := body["openid"]
	if openidAny == nil {
		return u, "", errors.Errorf("get openid fail! err: %+v", body["errmsg"])
	}

	if openid, ok := openidAny.(string); ok {
		// 生成token。如果数据库里没有该用户，则在该数据库生成该user
		user, err := dao.FindUserByOpenID(ctx, openid)
		if err != nil {
			return u, "", errors.WithMessage(err, "find user err")
		}

		if user != nil {
			return *user, user.CreateToken(""), nil
		}

		// 未注册过，需要先注册
		user = &bdm.User{
			OpenId:      openid,
			Gender:      gender,
			Username:    username,
			ClubId:      0,
			Intro:       "",
			AdminType:   constant.AdminType_NotAdmin,
			UserType:    0,
			Price:       0,
			Commissions: 0,
			AvatarURL:   avatarURL,
		}
		err = dao.SaveUser(ctx, user)
		if err != nil {
			return u, "", errors.WithMessage(err, "register user fail")
		}
		// 目前不需要sessionKey，所以暂时只是保存空sessionKey
		// sessionKey作用：https://developers.weixin.qq.com/community/develop/doc/00088a409fc308b765475fa4351000
		return *user, user.CreateToken(""), nil

	} else {
		return u, "", errors.Errorf("openid 不为 string. openid:%+v", openid)
	}
}
