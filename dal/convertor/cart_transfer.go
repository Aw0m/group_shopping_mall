package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func CartRdmToBdm(c rdm.Cart) bdm.Cart {
	return bdm.Cart{
		CartId:       c.CartId,
		CustomerId:   c.CustomerId,
		CommodityId:  c.CommodityId,
		InitPrice:    c.InitPrice,
		CreateTime:   c.CreateTime,
		IsDeleted:    c.IsDeleted,
		CommodityNum: c.CommodityNum,
	}
}

func CartBdmToRdm(c bdm.Cart) rdm.Cart {
	return rdm.Cart{
		CartId:       c.CartId,
		CustomerId:   c.CustomerId,
		CommodityId:  c.CommodityId,
		InitPrice:    c.InitPrice,
		CreateTime:   c.CreateTime,
		IsDeleted:    c.IsDeleted,
		CommodityNum: c.CommodityNum,
	}
}
