package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func UpdateUserinfo(ctx *gin.Context) {
	req, err := checkUpdateUserinfoReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.UpdateUserinfoRsp)
	err = biz.UpdateUserinfo(ctx, req.Username, req.Gender, req.AvatarURL, req.PhoneNum, req.AddressId)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkUpdateUserinfoReq(ctx *gin.Context) (req *rdm.UpdateUserinfoReq, retErr error) {
	req = new(rdm.UpdateUserinfoReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
