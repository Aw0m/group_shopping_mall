package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func BatchGetSellerByIdList(ctx context.Context, db *gorm.DB, idList []int64) (retSellerList []bdm.Seller, retErr error) {
	if len(idList) == 0 {
		return nil, nil
	}

	var rdmSellerList []rdm.Seller
	res := db.WithContext(ctx).Where("seller_id in (?)", idList).Find(&rdmSellerList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table seller err! err:%s", res.Error.Error())
	}

	retSellerList = make([]bdm.Seller, 0, len(rdmSellerList))
	for _, s := range rdmSellerList {
		retSellerList = append(retSellerList, convertor.SellerRdmToBdm(s))
	}
	return retSellerList, nil
}
