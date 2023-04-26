package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func GetCategoryList(ctx *gin.Context) {
	_, err := checkGetCategoryListReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.GetCategoryListRsp)
	rsp.CategoryList, err = biz.GetCategoryList(ctx)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkGetCategoryListReq(ctx *gin.Context) (req *rdm.GetCategoryListReq, retErr error) {
	req = new(rdm.GetCategoryListReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
