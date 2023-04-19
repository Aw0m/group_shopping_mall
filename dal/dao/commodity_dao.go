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
