package admin

import (
	"encoding/json"
	"net/http"
	"xiaomishop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

// 后台首页页面
func (con MainController) Index(c *gin.Context) {
  // 1. 获取session中的userinfo
  session := sessions.Default(c)
  userinfo := session.Get("userinfo")

  // 2. 类型断言判断userinfo是否是string
  userinfoStr, ok := userinfo.(string)
  if ok {
    // 2.1 获取用户信息
    var userinfoStruct []models.Manager
    json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

    // 2.2 获取所有的权限
    accessList := []models.Access{}
    models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

    // 2.3 获取当前角色拥有的权限, 并将所有权限id存入map对象中
    roleAccess := []models.RoleAccess{}
    models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
    roleAccessMap := make(map[int]int)
    for _, v := range roleAccess {
      roleAccessMap[v.AccessId] = v.AccessId
    }

    // 2.4 遍历所有的权限, 判断当前权限id是否在角色权限的map对象中, 如果存在则给当前权限增加checked属性。
    for i := 0; i < len(accessList); i++ {
      if _, ok := roleAccessMap[accessList[i].Id]; ok {
        accessList[i].Checked = true
      }

      for j := 0; j < len(accessList[i].AccessItem); j++ {
        if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
          accessList[i].AccessItem[j].Checked = true
        }
      }
    }

    c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
      "username": userinfoStruct[0].Username,
      "accessList": accessList,
      "isSuper": userinfoStruct[0].IsSuper,
    })
  } else {
    c.Redirect(http.StatusFound, "admin/login/login.html")
  }
}

// 后台欢迎页面
func (con MainController) Welcome(c *gin.Context) {
  c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
