package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetUserCount(ctx *gin.Context) {
	_, err := checkGetUserCountReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetUserCountRsp)
	rsp.UserCount, err = biz.GetUserCount(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetUserCountReq(ctx *gin.Context) (req *rdm.GetUserCountReq, retErr error) {
	req = new(rdm.GetUserCountReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
