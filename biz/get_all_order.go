package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetAllOrder(ctx *gin.Context, statusList []int, pageSize, pageNum int) (orderList []bdm.Order, retErr error) {
	if pageNum <= 0 || pageSize <= 0 {
		return nil, nil
	}

	offset := pageSize * (pageNum - 1)
	limit := pageSize

	orderList, err := dao.GetAllOrdersWithOffset(ctx, client_db.GetDB(), statusList, limit, offset)
	if err != nil {
		retErr = errors.WithMessage(err, "GetOrdersByCustomerId from db err!")
		return
	}

	return orderList, nil
}
