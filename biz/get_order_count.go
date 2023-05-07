package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
)

func GetOrderCount(ctx *gin.Context, statusList []int) (orderCount int64, retErr error) {
	orderCount, retErr = dao.GetOrderCount(ctx, client_db.GetDB(), statusList)
	if retErr != nil {
		retErr = errors.WithMessage(retErr, "GetOrderCount from db err!")
		return
	}
	return
}
