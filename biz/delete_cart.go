package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/utils/middleware"
)

func DeleteCart(ctx *gin.Context, cartId int64) (retErr error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		err = errors.WithMessage(err, "get id from token err")
		return err
	}

	cartList, err := dao.BatchGetCartById(ctx, client_db.GetDB(), []int64{cartId}, false)
	if err != nil {
		err = errors.WithMessage(err, "get cart from db fail")
		return err
	}
	if len(cartList) == 0 {
		err = errors.New("cart list is invalid")
		return err
	}

	cart := cartList[0]
	if cart.CustomerId != userId {
		err = errors.Errorf("custormerId does not mattch! customerId:%d, userId:%d", cart.CustomerId, userId)
		return err
	}

	return dao.UpdateCart(ctx, client_db.GetDB(), cartId, map[string]any{
		"is_deleted": true,
	})
}
