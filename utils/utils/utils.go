package utils

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = jsoniter.Unmarshal(body, &result)
	}

	return result, err
}

func HttpResp(ctx *gin.Context, rsp any, httpErrCode int, err error) {
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"rsp":     rsp,
		})
		return
	}
	ctx.JSON(httpErrCode, gin.H{
		"message": err.Error(),
		"rsp":     nil,
	})
}

func EncodeToMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	//将[]byte转成16进制
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func JoinInt(nums []int, sep string) string {
	if len(nums) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(strconv.Itoa(nums[0]))
	for i := 1; i < len(nums); i++ {
		b.WriteString(fmt.Sprintf("%s%d", sep, nums[i]))
	}
	return b.String()
}

func GetConfig[T any](path string) T {
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config T
	err = jsoniter.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func GetUserIdList[T any](objList []T, getIdFunc func(T) int64) []int64 {
	userIdSet := make(map[int64]struct{})
	for _, obj := range objList {
		userIdSet[getIdFunc(obj)] = struct{}{}
	}

	userIdList := make([]int64, 0, len(userIdSet))
	for userId := range userIdSet {
		userIdList = append(userIdList, userId)
	}

	return userIdList
}
