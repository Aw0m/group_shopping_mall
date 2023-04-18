package address

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetAddressList(ctx *gin.Context) {
	_, err := checkGetAddressListReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetAddressListRsp)
	rsp.AddressList, err = biz.GetAddressList(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetAddressListReq(ctx *gin.Context) (req *rdm.GetAddressListReq, retErr error) {
	req = new(rdm.GetAddressListReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
