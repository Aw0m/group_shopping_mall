package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetOrderDetail(ctx *gin.Context, orderId int64) (order bdm.Order, retErr error) {
	orderList, err := dao.BatchGetOrdersByOrderId(ctx, client_db.GetDB(), []int64{orderId}, true)
	if err != nil {
		retErr = errors.WithMessage(err, "get order from db fail")
		return
	}
	if len(orderList) == 0 {
		retErr = errors.Errorf("can not find order from db! OrderID:%d", orderId)
		return
	}

	order = orderList[0]

	return order, nil
}
