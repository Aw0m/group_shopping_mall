package convertor

import (
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

func UserRdmToBdm(u rdm.User) bdm.User {
	return bdm.User{
		UserId:      u.UserId,
		OpenId:      u.OpenId,
		Gender:      u.Gender,
		Username:    u.Username,
		ClubId:      u.ClubId,
		Intro:       u.Intro,
		AdminType:   u.AdminType,
		UserType:    u.UserType,
		Price:       u.Price,
		Commissions: u.Commissions,
		AvatarURL:   u.AvatarURL,
	}
}

func UserBdmToRdm(u bdm.User) rdm.User {
	return rdm.User{
		UserId:      u.UserId,
		OpenId:      u.OpenId,
		Gender:      u.Gender,
		Username:    u.Username,
		ClubId:      u.ClubId,
		Intro:       u.Intro,
		AdminType:   u.AdminType,
		UserType:    u.UserType,
		Price:       u.Price,
		Commissions: u.Commissions,
		AvatarURL:   u.AvatarURL,
	}
}
