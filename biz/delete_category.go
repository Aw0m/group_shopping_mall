package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
)

func DeleteCategory(ctx *gin.Context, categoryId int64) (retErr error) {
	retErr = dao.UpdateCategory(ctx, client_db.GetDB(), categoryId, map[string]interface{}{"is_deleted": 1})
	if retErr != nil {
		retErr = errors.WithMessage(retErr, "update category fail!")
		return retErr
	}
	return nil
}
