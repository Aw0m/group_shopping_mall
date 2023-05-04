package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/constant"
)

func ShipOrder(ctx *gin.Context, orderId int64, deliverymanName string, deliverymanPhoneNum string) (retErr error) {
	tx := client_db.GetDB().Begin()
	defer func() {
		if retErr != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// 1. 根据订单id获取订单信息
	orderList, err := dao.BatchGetOrdersByOrderId(ctx, tx, []int64{orderId}, false)
	if err != nil {
		retErr = errors.WithMessage(err, "BatchGetOrdersByOrderId from db err!")
		return retErr
	}
	if len(orderList) == 0 {
		retErr = errors.New("order not exist")
		return retErr
	}
	order := orderList[0]

	// 2. 校验订单信息
	if order.Status != constant.OrderStatus_Unshipped {
		retErr = errors.Errorf("order status not match! order_id:%d, order_status:%d", order.OrderId, order.Status)
		return retErr
	}

	// 3. 更新订单状态
	err = dao.UpdateOrder(ctx, tx, order.OrderId, map[string]any{
		"status":                constant.OrderStatus_Shipped,
		"deliveryman_name":      deliverymanName,
		"deliveryman_phone_num": deliverymanPhoneNum,
	})
	if err != nil {
		retErr = errors.WithMessage(err, "update order status fail")
		return retErr
	}

	return
}
