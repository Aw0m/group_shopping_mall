package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/middleware"
)

func UpdateCartNum(ctx *gin.Context, cartId int64, quantity int) (retErr error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		retErr = errors.WithMessage(err, "get user id failed from token")
		return
	}

	// 拉取cart信息并校验
	cartList, err := dao.BatchGetCartById(ctx, client_db.GetDB(), []int64{cartId}, false)
	if err != nil {
		retErr = errors.WithMessage(err, "get cart from db failed")
		return
	}
	if len(cartList) == 0 {
		retErr = errors.WithMessage(err, "cart not exist")
		return
	}
	cart := cartList[0]
	if userId != cart.CustomerId {
		retErr = errors.WithMessage(err, "cart not belong to this user")
		return
	}

	// 更新cart数量
	err = dao.UpdateCart(ctx, client_db.GetDB(), cartId, map[string]any{"commodity_num": quantity})
	if err != nil {
		retErr = errors.WithMessage(err, "update cart failed")
		return
	}
	return
}
