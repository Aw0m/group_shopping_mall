package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/utils/middleware"
)

func GetUserinfo(ctx *gin.Context) (bdm.User, error) {
	_, userId, err := middleware.GetID(ctx)
	if err != nil {
		return bdm.User{}, errors.WithMessage(err, "get id from context failed")
	}

	user, err := dao.FindUserByUserId(ctx, userId)
	if err != nil {
		return bdm.User{}, errors.WithMessage(err, "find user from db failed")
	}
	if user == nil {
		return bdm.User{}, errors.Errorf("user not found! user_id:%d", userId)
	}

	return *user, nil
}
