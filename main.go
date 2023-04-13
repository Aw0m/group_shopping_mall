package main

import (
	"flag"
	"fmt"

	"group_shopping_mall/utils/constant"

	"github.com/gin-gonic/gin"

	"group_shopping_mall/controller/user"
	"group_shopping_mall/dal/client_db"
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
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
