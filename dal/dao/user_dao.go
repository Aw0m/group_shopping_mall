package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/convertor"
	"group_shopping_mall/model/bdm"
	"group_shopping_mall/model/rdm"
)

// FindUserByOpenID 通过id查找对应的聚合
func FindUserByOpenID(_ context.Context, openid string) (*bdm.User, error) {
	db := client_db.GetDB()
	userRdm := new(rdm.User)
	res := db.Where("open_id = ?", openid).First(userRdm)
	if res.Error == nil {
		u := convertor.UserRdmToBdm(*userRdm)
		return &u, nil
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, errors.Errorf("query from table user err! err:%s", res.Error.Error())
}

// SaveUser 保存user。由于函数中存在查后更新，可能存在并发问题
func SaveUser(ctx context.Context, userBdm *bdm.User) error {
	user := convertor.UserBdmToRdm(*userBdm)

	db := client_db.GetDB()
	u, err := FindUserByOpenID(ctx, user.OpenId)
	if err != nil {
		return errors.WithMessage(err, "find user by id err!")
	}

	// 已存在db中则直接update
	if u != nil {
		userMap := getUserMap(userBdm)
		db.Model(&rdm.User{}).Where("open_id = ?", user.OpenId).Updates(userMap)
		return nil
	}

	// 否则在db中create
	res := db.Create(&user)
	return res.Error
}

func BatchUpdateUsers(_ context.Context, db *gorm.DB, userIds []int64, userMap map[string]interface{}) error {
	res := db.Model(&rdm.User{}).Where("user_id IN ?", userIds).Updates(userMap)
	if res.Error != nil {
		return errors.Errorf("update user fail! err:%s", res.Error.Error())
	}
	return nil
}

// gorm更新0值的字段，总是失败，所以需要自己构造map
func getUserMap(user *bdm.User) map[string]interface{} {
	userMap := make(map[string]interface{})
	userMap["user_id"] = user.UserId
	userMap["open_id"] = user.OpenId
	userMap["gender"] = user.Gender
	userMap["username"] = user.Username
	userMap["avatar_url"] = user.AvatarUrl
	userMap["phone_num"] = user.PhoneNum
	userMap["address_id"] = user.AddressId
	return userMap
}

// FindUserByUserId 查询指定userId的User
func FindUserByUserId(_ context.Context, userId int64) (*bdm.User, error) {
	db := client_db.GetDB()
	userRdm := new(rdm.User)
	res := db.Where("user_id = ?", userId).First(userRdm)
	if res.Error == nil {
		u := convertor.UserRdmToBdm(*userRdm)
		return &u, nil
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, errors.Errorf("query from table user err! err:%s", res.Error.Error())
}

func GetUserWithOffset(ctx context.Context, db *gorm.DB, offset, limit int) ([]bdm.User, error) {
	userRdms := make([]rdm.User, 0)
	res := db.Model(&rdm.User{}).Offset(offset).Limit(limit).Find(&userRdms)
	if res.Error != nil {
		return nil, errors.Errorf("query from table user err! err:%s", res.Error.Error())
	}

	users := make([]bdm.User, 0)
	for _, userRdm := range userRdms {
		users = append(users, convertor.UserRdmToBdm(userRdm))
	}
	return users, nil
}

func GetUserCount(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	res := db.Model(&rdm.User{}).Count(&count)
	if res.Error != nil {
		return 0, errors.Errorf("query from table user err! err:%s", res.Error.Error())
	}
	return count, nil
}
