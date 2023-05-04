package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/middleware"
)

func GetOrderList(ctx *gin.Context, statusList []int) (orderList []bdm.Order, retErr error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		retErr = errors.WithMessage(err, "get user id from token fail")
		return
	}

	orderList, err = dao.GetOrdersByCustomerId(ctx, client_db.GetDB(), userId, statusList)
	if err != nil {
		retErr = errors.WithMessage(err, "GetOrdersByCustomerId from db err!")
		return
	}

	return orderList, nil
}
