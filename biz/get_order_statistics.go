package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/middleware"
)

func GetOrderStatistics(ctx *gin.Context) (pendingNum, unshippedNum, shippedNum, completedNum int, retErr error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		retErr = errors.WithMessage(err, "get user id from token fail")
		return
	}

	// 1. 根据用户id获取订单信息
	orderList, err := dao.GetOrdersByCustomerId(
		ctx,
		client_db.GetDB(),
		userId,
		[]int{constant.OrderStatus_Pending, constant.OrderStatus_Shipped, constant.OrderStatus_Unshipped, constant.OrderStatus_Completed},
		false,
	)
	if err != nil {
		retErr = errors.WithMessage(err, "GetOrdersByCustomerId from db err!")
		return
	}

	// 2. 统计订单状态数量
	for _, order := range orderList {
		switch order.Status {
		case constant.OrderStatus_Pending:
			pendingNum++
		case constant.OrderStatus_Unshipped:
			unshippedNum++
		case constant.OrderStatus_Shipped:
			shippedNum++
		case constant.OrderStatus_Completed:
			completedNum++
		}
	}

	return
}
