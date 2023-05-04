package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/pkg/errors"

	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func InsertOrders(_ context.Context, db *gorm.DB, ordersList []bdm.Order) error {
	var rdmList []rdm.Order
	for _, o := range ordersList {
		rdmList = append(rdmList, convertor.OrderBdmToRdm(o))
	}
	res := db.Create(&rdmList)
	if res != nil && res.Error != nil {
		return errors.Errorf("insert orders fail! err:%s", res.Error.Error())
	}
	return nil
}
