package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/middleware"
)

func UpdateUserinfo(ctx *gin.Context, username, gender, avatarUrl, phoneNum string, addressId int64) (err error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		return errors.WithMessage(err, "get id from token failed")
	}

	updateMap := make(map[string]interface{})
	if username != "" {
		updateMap["username"] = username
	}

	if gender != "" {
		updateMap["gender"] = gender
	}

	if avatarUrl != "" {
		updateMap["avatar_url"] = avatarUrl
	}

	if phoneNum != "" {
		updateMap["phone_num"] = phoneNum
	}

	if addressId != 0 {
		// 检查addressId是否存在
		addressList, err := GetAddressList(ctx)
		if err != nil {
			return errors.WithMessage(err, "get addressList failed")
		}
		for _, address := range addressList {
			if address.AddressId == addressId {
				// 此时addressId存在
				updateMap["address_id"] = addressId
				break
			}
		}
	}

	err = dao.BatchUpdateUsers(ctx, client_db.GetDB(), []int64{userId}, updateMap)
	if err != nil {
		return errors.WithMessage(err, "update user failed")
	}
	return nil
}
