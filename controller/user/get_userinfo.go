package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetUserinfo(ctx *gin.Context) {
	_, err := checkGetUserinfoReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetUserinfoRsp)
	rsp.User, err = biz.GetUserinfo(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetUserinfoReq(ctx *gin.Context) (req *rdm.GetUserinfoReq, retErr error) {
	req = new(rdm.GetUserinfoReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
