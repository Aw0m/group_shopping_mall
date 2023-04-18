package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/middleware"
)

func GetCartList(ctx *gin.Context) ([]bdm.Cart, error) {
	_, userid, err := middleware.GetID(ctx)
	if err != nil {
		err = errors.WithMessage(err, "GetID from token err!")
	}

	cartList, err := dao.GetCartListByUserId(ctx, client_db.GetDB(), userid, false)
	if err != nil {
		err = errors.WithMessage(err, "GetCartListByUserId from db err!")
		return nil, err
	}

	return cartList, nil
}
