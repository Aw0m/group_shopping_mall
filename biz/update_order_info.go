package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
)

func UpdateOrderInfo(ctx *gin.Context, OrderId int64,
	sellerName, commodityName, deliverymanName, deliverymanPhoneNum, phoneNum, consigneeName, addressDetailInfo string,
	price float64, quantity, status int, isDeleted bool) (retErr error) {

	updateMap := map[string]any{
		"seller_name":           sellerName,
		"commodity_name":        commodityName,
		"deliveryman_name":      deliverymanName,
		"deliveryman_phone_num": deliverymanPhoneNum,
		"phone_num":             phoneNum,
		"consignee_name":        consigneeName,
		"address_detail_info":   addressDetailInfo,
		"price":                 price,
		"quantity":              quantity,
		"status":                status,
		"is_deleted":            isDeleted,
	}
	err := dao.UpdateOrder(ctx, client_db.GetDB(), OrderId, updateMap)
	if err != nil {
		retErr = errors.WithMessage(err, "update order info fail")
		return
	}
	return nil
}
