package biz

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_redis"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetAddressList(ctx *gin.Context) (addressList []bdm.ShoppingAddress, err error) {
	key := dao.GetAddressKey()
	val, err := client_redis.GetValue(ctx, key)
	if err != nil {
		err = errors.WithMessage(err, "get addressVal from redis failed")
		return nil, err
	}

	if len(val) == 0 {
		return []bdm.ShoppingAddress{}, nil
	}
	err = jsoniter.UnmarshalFromString(val, &addressList)
	if err != nil {
		err = errors.WithMessage(err, "unmarshal addressList failed")
		return nil, err
	}

	return addressList, nil
}
