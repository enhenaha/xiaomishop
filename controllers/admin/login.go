package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xiaomishop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
  BaseController
}

// 登录页
func (con LoginController) Index(c *gin.Context) {
  c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

// 登录接口
func (con LoginController) DoLogin(c *gin.Context) {
  username := c.PostForm("username")
  password := c.PostForm("password")
  vaptchaId := c.PostForm("captchaId")
  verifyValue := c.PostForm("verifyValue")

  // 1. 验证验证码
  if models.VerifyCaptcha(vaptchaId, verifyValue) {
    // 2. 验证账号密码
    userinfoList := []models.Manager{}
    password := models.Md5(password) // 对密码进行md5加密
    models.DB.Where("username = ? AND password = ?", username, password).Find(&userinfoList)

    if len(userinfoList) > 0 {
      // 3. 执行登录, 保存用户信息并跳转
      session := sessions.Default(c)
      userinfoSlice, _ := json.Marshal(userinfoList) // 将切片转为JSON字符串
      session.Set("userinfo", string(userinfoSlice))
      session.Save()

      con.Success(c, "验证码验证成功", "/admin")
    } else {
      con.Error(c, "用户名或密码错误", "/admin/login")
    }
    
  } else {
    con.Error(c, "验证码验证失败", "/admin/login")
  }
}

// 图形验证码接口
func (con LoginController) Captcha(c *gin.Context) {
  id, b64s, err := models.MakeCaptcha()

  if err != nil {
    fmt.Println("err: ", err)
  }

  c.JSON(http.StatusOK, gin.H{
    "captchaId": id,
    "captchaImage": b64s,
  })
}

// 退出登录
func (con LoginController) LoginOut(c *gin.Context) {
  session := sessions.Default(c)
  session.Delete("userinfo")
  session.Save()
  con.Success(c, "退出登录成功", "/admin/login")
}
