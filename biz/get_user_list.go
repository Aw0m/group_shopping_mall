package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetUserList(ctx *gin.Context, pageNum, pageSize int) (userList []bdm.User, retErr error) {
	if pageNum <= 0 || pageSize <= 0 {
		return nil, nil
	}

	offset := pageSize * (pageNum - 1)
	limit := pageSize

	userList, retErr = dao.GetUserWithOffset(ctx, client_db.GetDB(), offset, limit)
	if retErr != nil {
		retErr = errors.WithMessage(retErr, "dao.GetUserWithOffset failed")
		return nil, retErr
	}
	return
}
