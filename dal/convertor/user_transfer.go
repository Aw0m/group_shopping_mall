package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func UserRdmToBdm(u rdm.User) bdm.User {
	return bdm.User{
		UserId:    u.UserId,
		OpenId:    u.OpenId,
		Gender:    u.Gender,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl,
		PhoneNum:  u.PhoneNum,
		AddressId: u.AddressId,
	}
}

func UserBdmToRdm(u bdm.User) rdm.User {
	return rdm.User{
		UserId:    u.UserId,
		OpenId:    u.OpenId,
		Gender:    u.Gender,
		Username:  u.Username,
		AvatarUrl: u.AvatarUrl,
		PhoneNum:  u.PhoneNum,
		AddressId: u.AddressId,
	}
}
