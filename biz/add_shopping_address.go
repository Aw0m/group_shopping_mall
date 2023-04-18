package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/utils"
)

func AddShoppingAddress(ctx *gin.Context, addressName, detailInfo string) (retErr error) {
	//TODO 需要加分布式锁
	addressList, err := GetAddressList(ctx)
	if err != nil {
		err = errors.WithMessage(err, "get addressList failed")
		return err
	}

	address := bdm.ShoppingAddress{
		AddressId:   utils.GetSnowflakeId(),
		AddressName: addressName,
		DetailInfo:  detailInfo,
	}
	addressList = append(addressList, address)

	err = dao.SaveShoppingAddress(ctx, addressList)
	if err != nil {
		err = errors.WithMessage(err, "save addressList failed")
		return err
	}

	return nil
}
