package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/dao"
	"group_shopping_mall/model/bdm"
)

func GetCategoryList(ctx *gin.Context) (categoryList []bdm.Category, retErr error) {
	categoryList, err := dao.GetCategoryList(ctx, client_db.GetDB())
	if err != nil {
		err = errors.WithMessage(err, "get category list from db err")
		return nil, err
	}

	return categoryList, nil
}
