package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetUserList(ctx *gin.Context) {
	req, err := checkGetUserListReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetUserListRsp)
	rsp.UserList, err = biz.GetUserList(ctx, req.PageNum, req.PageSize)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetUserListReq(ctx *gin.Context) (req *rdm.GetUserListReq, retErr error) {
	req = new(rdm.GetUserListReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
