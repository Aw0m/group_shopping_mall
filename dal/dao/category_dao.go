package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func GetCategoryList(ctx *gin.Context, db *gorm.DB) (categoryList []bdm.Category, retErr error) {
	categoryRdmList := make([]rdm.Category, 0)
	res := db.WithContext(ctx).Find(&categoryRdmList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table category err! err:%s", res.Error.Error())
	}

	categoryList = make([]bdm.Category, 0, len(categoryRdmList))
	for _, categoryRdm := range categoryRdmList {
		categoryList = append(categoryList, convertor.CategoryRdmToBdm(categoryRdm))
	}
	return categoryList, nil
}

func InsertCategory(ctx *gin.Context, db *gorm.DB, category bdm.Category) (retErr error) {
	categoryRdm := convertor.CategoryBdmToRdm(category)
	res := db.WithContext(ctx).Create(&categoryRdm)
	if res.Error != nil {
		return errors.Errorf("create category err! err:%s", res.Error.Error())
	}
	return nil
}
