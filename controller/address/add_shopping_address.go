package address

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func AddShoppingAddress(ctx *gin.Context) {
	req, err := checkAddShoppingAddressReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.AddShoppingAddressRsp)
	err = biz.AddShoppingAddress(ctx, req.AddressName, req.DetailInfo)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkAddShoppingAddressReq(ctx *gin.Context) (req *rdm.AddShoppingAddressReq, retErr error) {
	req = new(rdm.AddShoppingAddressReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
