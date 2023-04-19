package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func CommodityBdmToRdm(c bdm.Commodity) rdm.Commodity {
	return rdm.Commodity{
		CommodityId:     c.CommodityId,
		CommodityName:   c.CommodityName,
		Price:           c.Price,
		SellerId:        c.SellerId,
		CategoryId:      c.CategoryId,
		Content:         c.Content,
		AddressList:     c.AddressList,
		CommodityStatus: c.CommodityStatus,
		IsDeleted:       c.IsDeleted,
	}
}

func CommodityRdmToBdm(c rdm.Commodity) bdm.Commodity {
	return bdm.Commodity{
		CommodityId:     c.CommodityId,
		CommodityName:   c.CommodityName,
		Price:           c.Price,
		SellerId:        c.SellerId,
		CategoryId:      c.CategoryId,
		Content:         c.Content,
		AddressList:     c.AddressList,
		CommodityStatus: c.CommodityStatus,
		IsDeleted:       c.IsDeleted,
	}
}
