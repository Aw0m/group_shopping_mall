package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func BatchGetCommodityByIdList(ctx context.Context, db *gorm.DB, idList []int64) (retCommodityList []bdm.Commodity, retErr error) {
	if len(idList) == 0 {
		return nil, nil
	}

	var rdmCommodityList []rdm.Commodity
	res := db.WithContext(ctx).Where("commodity_id in (?)", idList).Find(&rdmCommodityList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table commodity err! err:%s", res.Error.Error())
	}

	retCommodityList = make([]bdm.Commodity, 0, len(rdmCommodityList))
	for _, c := range rdmCommodityList {
		retCommodityList = append(retCommodityList, convertor.CommodityRdmToBdm(c))
	}
	return retCommodityList, nil
}

func GetCommodityFromCategory(ctx context.Context, db *gorm.DB, categoryId int64, limit, offset int, withDeleted bool) (categoryList []bdm.Commodity, retErr error) {
	var rdmCategoryList []rdm.Commodity
	if !withDeleted {
		db = db.Where("is_deleted = ?", false)
	}

	// categoryId == 0 means get all commodities
	if categoryId > 0 {
		db = db.Where("category_id = ?", categoryId)
	}

	res := db.WithContext(ctx).
		Order("commodity_id DESC").
		Limit(limit).
		Offset(offset).
		Find(&rdmCategoryList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table category err! err:%s", res.Error.Error())
	}

	categoryList = make([]bdm.Commodity, 0, len(rdmCategoryList))
	for _, c := range rdmCategoryList {
		categoryList = append(categoryList, convertor.CommodityRdmToBdm(c))
	}
	return categoryList, nil
}

func GetCommodityCount(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	res := db.Model(&rdm.Commodity{}).Count(&count)
	if res.Error != nil {
		return 0, errors.Errorf("query from table user err! err:%s", res.Error.Error())
	}
	return count, nil
}

func UpdateCommodity(_ context.Context, db *gorm.DB, commodityId int64, updateMap map[string]any) error {
	res := db.Model(&rdm.Commodity{}).Where("commodity_id = ?", commodityId).Updates(updateMap)
	if res.Error != nil {
		return errors.Errorf("update commodity fail! err:%s", res.Error.Error())
	}
	return nil
}
