package main

import (
	"text/template"
	"xiaomishop/models"
	"xiaomishop/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
  // 1. 创建一个默认的路由引擎
  r := gin.Default()

  // 2. 自定义模板函数  注意要把这个函数放在加载模板前
  r.SetFuncMap(template.FuncMap{
    "UnixToTime": models.UnixToTime,
  })

  // 3. 加载模板
  r.LoadHTMLGlob("./templates/**/**/*")

  // 4. 配置静态web服务
  r.Static("/static", "./static")

  // 5. 配置session中间件
  store := cookie.NewStore([]byte("secret111")) // "secret111"参数是用于加密的密钥
  r.Use(sessions.Sessions("mysession", store)) // "mysession"是设置在客户端浏览器cookie的属性名

  // 6. 配置路由
  routers.AdminRoutersInit(r)

  // 7. 启动一个web服务
  r.Run(":8000")
}