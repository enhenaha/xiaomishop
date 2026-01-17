package admin

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type BaseController struct{}

// 成功页面
func (con BaseController) Success(c *gin.Context, message, redirectUrl string) {
  c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
    "message": message,
    "redirectUrl": redirectUrl,
  })
}

// 失败页面
func (con BaseController) Error(c *gin.Context, message, redirectUrl string) {
  c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
    "message": message,
    "redirectUrl": redirectUrl,
  })
}