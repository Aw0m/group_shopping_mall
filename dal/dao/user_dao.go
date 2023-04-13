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

func UpdateUser(_ context.Context, db *gorm.DB, user *bdm.User) error {
	userMap := getUserMap(user)
	res := db.Model(&rdm.User{}).Where("user_id = ?", user.UserId).Updates(userMap)
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
	userMap["club_id"] = user.ClubId
	userMap["intro"] = user.Intro
	userMap["admin_type"] = user.AdminType
	userMap["user_type"] = user.UserType
	userMap["price"] = user.Price
	userMap["commissions"] = user.Commissions
	userMap["avatar_url"] = user.AvatarURL
	return userMap
}

// FindUserByClubIdWithOffset 分页查询指定clubId的User
func FindUserByClubIdWithOffset(_ context.Context, clubId int64, limit, offset int, keyword string) (userList []bdm.User, retErr error) {
	db := client_db.GetDB()
	var rdmList []rdm.User
	res := db.Where("club_id = ? AND (user_id = ? OR username LIKE ?)", clubId, keyword, "%"+keyword+"%").Limit(limit).Offset(offset).Find(&rdmList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table user fail! err:%s", res.Error.Error())
	}

	for _, u := range rdmList {
		userList = append(userList, convertor.UserRdmToBdm(u))
	}
	return userList, nil
}

func FindUserByClubId(_ context.Context, db *gorm.DB, clubId int64) ([]bdm.User, error) {
	var rdmList []rdm.User
	res := db.Model(&rdm.User{}).Where("club_id = ?", clubId).Find(&rdmList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table user fail! err:%s", res.Error.Error())
	}

	var userList []bdm.User
	for _, u := range rdmList {
		userList = append(userList, convertor.UserRdmToBdm(u))
	}
	return userList, nil
}

// FindUserByClubIdAndUserId 查询指定clubId和userId的User
func FindUserByClubIdAndUserId(_ context.Context, clubId, userId int64) (*bdm.User, error) {
	db := client_db.GetDB()
	userRdm := new(rdm.User)
	res := db.Where("club_id = ? AND user_id = ?", clubId, userId).First(userRdm)
	if res.Error == nil {
		u := convertor.UserRdmToBdm(*userRdm)
		return &u, nil
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, errors.Errorf("query from table user err! err:%s", res.Error.Error())
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

// BatchGetUserByIdList 批量查询指定id的user
func BatchGetUserByIdList(_ context.Context, idList []int64) (userList []bdm.User, retErr error) {
	db := client_db.GetDB()
	var rdmList []rdm.User
	res := db.Where("user_id IN ?", idList).Find(&rdmList)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Errorf("query from table user fail! err:%s", res.Error.Error())
	}

	for _, u := range rdmList {
		userList = append(userList, convertor.UserRdmToBdm(u))
	}
	return userList, nil
}
