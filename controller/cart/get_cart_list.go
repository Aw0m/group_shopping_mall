package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCartList(ctx *gin.Context) {
	_, err := checkGetCartListReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCartListRsp)
	rsp.CommodityItemList, err = biz.GetCartList(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCartListReq(ctx *gin.Context) (req *rdm.GetCartListReq, retErr error) {
	req = new(rdm.GetCartListReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
