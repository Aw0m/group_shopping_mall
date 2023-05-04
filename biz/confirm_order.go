package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/middleware"
)

func ConfirmOrder(ctx *gin.Context, orderId int64) (retErr error) {
	tx := client_db.GetDB().Begin()
	defer func() {
		if retErr != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		return errors.WithMessage(err, "get user id fail")
	}
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
	if order.CustomerId != userId {
		retErr = errors.Errorf("user is not the customer of order.userID:%d, order customerID:%d", userId, orderId)
		return retErr
	}

	if order.Status != constant.OrderStatus_Pending {
		retErr = errors.Errorf("order status not match! order_id:%d, order_status:%d", order.OrderId, order.Status)
		return retErr
	}

	// 3. 更新订单状态
	err = dao.UpdateOrderStatus(ctx, tx, order.OrderId, constant.OrderStatus_Unshipped)
	if err != nil {
		retErr = errors.WithMessage(err, "update order status fail")
		return retErr
	}

	return nil
}
