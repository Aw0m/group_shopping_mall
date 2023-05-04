package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

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

func BatchGetOrdersByOrderId(_ context.Context, db *gorm.DB, orderIdList []int64, withDeleted bool) ([]bdm.Order, error) {
	var orderList []rdm.Order
	if !withDeleted {
		db = db.Where("is_deleted = ?", false)
	}
	res := db.Where("order_id in ?", orderIdList).Find(&orderList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("find order by order_id fail! err:%s", res.Error.Error())
	}

	var bdmOrderList []bdm.Order
	for _, order := range orderList {
		bdmOrderList = append(bdmOrderList, convertor.OrderRdmToBdm(order))
	}
	return bdmOrderList, nil
}

func UpdateOrderStatus(_ context.Context, db *gorm.DB, orderId int64, status int) error {
	res := db.Model(&rdm.Order{}).Where("order_id = ?", orderId).Update("status", status)
	if res.Error != nil {
		return errors.Errorf("update order status fail! err:%s", res.Error.Error())
	}
	return nil
}

func UpdateOrder(_ context.Context, db *gorm.DB, orderId int64, updateMap map[string]any) error {
	res := db.Model(&rdm.Order{}).Where("order_id = ?", orderId).Updates(updateMap)
	if res.Error != nil {
		return errors.Errorf("update order fail! err:%s", res.Error.Error())
	}
	return nil
}
