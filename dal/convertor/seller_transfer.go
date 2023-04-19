package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func SellerRdmToBdm(s rdm.Seller) bdm.Seller {
	return bdm.Seller{
		SellerId:   s.SellerId,
		SellerName: s.SellerName,
		PhoneNum:   s.PhoneNum,
		Status:     s.Status,
		CreateTime: s.CreateTime,
		IsDeleted:  s.IsDeleted,
	}
}

func SellerBdmToRdm(s bdm.Seller) rdm.Seller {
	return rdm.Seller{
		SellerId:   s.SellerId,
		SellerName: s.SellerName,
		PhoneNum:   s.PhoneNum,
		Status:     s.Status,
		CreateTime: s.CreateTime,
		IsDeleted:  s.IsDeleted,
	}
}
