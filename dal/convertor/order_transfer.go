package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func OrderBdmToRdm(order bdm.Order) rdm.Order {
	return rdm.Order{
		OrderId:             order.OrderId,
		SellerId:            order.SellerId,
		SellerName:          order.SellerName,
		CommodityId:         order.CommodityId,
		CommodityName:       order.CommodityName,
		CommodityImageURL:   order.CommodityImageURL,
		CustomerId:          order.CustomerId,
		DeliverymanName:     order.DeliverymanName,
		DeliverymanPhoneNum: order.DeliverymanPhoneNum,
		PhoneNum:            order.PhoneNum,
		ConsigneeName:       order.ConsigneeName,
		AddressId:           order.AddressId,
		AddressDetailInfo:   order.AddressDetailInfo,
		Price:               order.Price,
		Quantity:            order.Quantity,
		Status:              order.Status,
		CreateTime:          order.CreateTime,
		UpdateTime:          order.UpdateTime,
		IsDeleted:           order.IsDeleted,
	}
}

func OrderRdmToBdm(order rdm.Order) bdm.Order {
	return bdm.Order{
		OrderId:             order.OrderId,
		SellerId:            order.SellerId,
		SellerName:          order.SellerName,
		CommodityId:         order.CommodityId,
		CommodityName:       order.CommodityName,
		CommodityImageURL:   order.CommodityImageURL,
		CustomerId:          order.CustomerId,
		DeliverymanName:     order.DeliverymanName,
		DeliverymanPhoneNum: order.DeliverymanPhoneNum,
		PhoneNum:            order.PhoneNum,
		ConsigneeName:       order.ConsigneeName,
		AddressId:           order.AddressId,
		AddressDetailInfo:   order.AddressDetailInfo,
		Price:               order.Price,
		Quantity:            order.Quantity,
		Status:              order.Status,
		CreateTime:          order.CreateTime,
		UpdateTime:          order.UpdateTime,
		IsDeleted:           order.IsDeleted,
	}
}
