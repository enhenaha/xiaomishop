package middlewars

import (
	// "encoding/json"
	"encoding/json"
	"xiaomishop/models"

	// "net/http"
	// "xiaomishop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {
  // 1. 获取url中的path
  pathname := c.Request.URL.Path

  // 2. 获取session里面保存的用户信息
  session := sessions.Default(c)
  userinfo := session.Get("userinfo") // 返回一个空接口类型, 如果获取不到会返回nil
  userinfoStr, ok := userinfo.(string) // 类型断言判断userinfo是一个string类型

  // 3. 判断用户信息类型是否正确
  if ok {
    // 3.1 判断userinfo里面的信息是否存在
    var userinfoStruct []models.Manager // 实例化结构体
    err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

    // 有错误 或 用户信息没有用户名
    if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
      // 用户信息没有用户名 且 跳转到需要权限的页面
      if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
        c.Redirect(302, "/admin/login")
      }
    } else { // 用户登录成功, 进行权限判断
      // 1. 获取当前用户的角色拥有的权限, 并将所有权限id存入map对象中
      roleAccess := []models.RoleAccess{}
      models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
      roleAccessMap := make(map[int]int)
      for _, v := range roleAccess {
        roleAccessMap[v.AccessId] = v.AccessId
      }

      // 2. 获取当前访问的url对应的权限id, 判断该权限id是否在角色权限的map对象中。


    }

  } else {
    // 用户没有登录, 跳转到登录页 (排除不需要做权限判断的路由, 避免陷入死循环)
    if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
      c.Redirect(302, "/admin/login")
    }
  }
}