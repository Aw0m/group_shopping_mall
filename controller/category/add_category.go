package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/biz"
	"group_shopping_mall/model/rdm"
	"group_shopping_mall/utils/utils"
)

func AddCategory(ctx *gin.Context) {
	req, err := checkAddCategoryReq(ctx)
	if err != nil {
		utils.HttpResp(ctx, nil, http.StatusBadRequest, err)
		return
	}

	rsp := new(rdm.AddCategoryRsp)
	err = biz.AddCategory(ctx, req.CategoryName)
	utils.HttpResp(ctx, rsp, http.StatusInternalServerError, err)
	return
}

func checkAddCategoryReq(ctx *gin.Context) (req *rdm.AddCategoryReq, retErr error) {
	req = new(rdm.AddCategoryReq)
	err := utils.ReqBindAndCheck(ctx, req)
	if err != nil {
		return req, err
	}

	return req, nil
}
