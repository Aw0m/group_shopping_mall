package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func AddCategory(ctx *gin.Context, categoryName string) (retErr error) {
	newCategory := bdm.Category{
		CategoryName: categoryName,
		IsDeleted:    false,
	}
	retErr = dao.InsertCategory(ctx, client_db.GetDB(), newCategory)
	if retErr != nil {
		retErr = errors.WithMessage(retErr, "insert category err!")
		return retErr
	}

	return nil
}
