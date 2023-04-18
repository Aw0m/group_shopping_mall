package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/controller/address"
	"group_shopping_mall/controller/user"
	"group_shopping_mall/dal/client_db"
	"group_shopping_mall/dal/client_redis"
	"group_shopping_mall/utils/constant"
	"group_shopping_mall/utils/middleware"
	"group_shopping_mall/utils/utils"
)

var (
	port int
)

func Init() {
	flag.IntVar(&port, "port", 8080, "the port to start gin")
	flag.Parse()

	utils.ValidateInit()
	utils.InitSnowflake()
	constant.InitConstant("constantConfig.json")
	client_db.InitDB("dbConfig.json")
	client_redis.InitRedis()
}
func main() {
	Init()

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", user.WXLogin)

	// 用户
	userGroup := r.Group("/user")
	userGroup.Use(middleware.Authorize())
	{
		userGroup.POST("/get_userinfo", user.GetUserinfo)
		userGroup.POST("/update_userinfo", user.UpdateUserinfo)
	}

	// 自提点
	addressGroup := r.Group("/address")
	addressGroup.Use(middleware.Authorize())
	{
		addressGroup.POST("/get_address_list", address.GetAddressList)
		addressGroup.POST("/add_shopping_address", address.AddShoppingAddress)
	}

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
