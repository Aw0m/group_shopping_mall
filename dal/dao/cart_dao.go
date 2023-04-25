package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

// GetCartListByUserId 获取userId对应的购物车列表
func GetCartListByUserId(_ context.Context, db *gorm.DB, userId int64, withDelete bool) ([]bdm.Cart, error) {
	cartRdmList := make([]rdm.Cart, 0)
	txTemp := db.Where("customer_id = ?", userId)
	if !withDelete {
		txTemp = txTemp.Where("is_deleted = ?", false)
	}
	res := txTemp.Order("-create_time").Find(&cartRdmList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table cart err! err:%s", res.Error.Error())
	}

	cartBdmList := make([]bdm.Cart, 0, len(cartRdmList))
	for _, cartRdm := range cartRdmList {
		cartBdmList = append(cartBdmList, convertor.CartRdmToBdm(cartRdm))
	}
	return cartBdmList, nil
}

// UpdateCart 更新Cart
func UpdateCart(_ context.Context, db *gorm.DB, cartId int64, updateMap map[string]any) error {
	res := db.Model(&rdm.Cart{}).Where("cart_id = ?", cartId).Updates(updateMap)
	if res.Error != nil {
		return errors.Errorf("update cart fail! err:%s", res.Error.Error())
	}
	return nil
}

// CreateCart 创建Cart
func CreateCart(_ context.Context, db *gorm.DB, cart bdm.Cart) error {
	cartRdm := convertor.CartBdmToRdm(cart)
	res := db.Create(&cartRdm)
	if res.Error != nil {
		return errors.Errorf("create cart fail! err:%s", res.Error.Error())
	}
	return nil
}
