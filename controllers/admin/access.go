package admin

import (
	"net/http"
	"strings"
	"xiaomishop/models"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
  BaseController
}

// 权限列表页面
func (con AccessController) Index(c *gin.Context) {
  accessList := []models.Access{}
  models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

  c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
    "accessList": accessList,
  })
}

// 增加权限页面
func (con AccessController) Add(c *gin.Context) {
  // 1. 获取顶级模块列表
  accessList := []models.Access{}
  models.DB.Where("module_id = ?", 0).Find(&accessList)

  c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
    "accessList": accessList,
  })
}

// 增加权限接口
func (con AccessController) DoAdd(c *gin.Context) {
  // 1. 获取数据
  moduleName := strings.Trim(c.PostForm("module_name"), " ")
  actionName := c.PostForm("action_name")
  accessType, err1 := models.Int(c.PostForm("type"))
  url := c.PostForm("url")
  moduleId, err2 := models.Int(c.PostForm("module_id"))
  sort, err3 := models.Int(c.PostForm("sort"))
  status, err4 := models.Int(c.PostForm("status"))
  description := c.PostForm("description")

  if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
    con.Error(c, "传入参数错误", "admin/access/add")
    return
  }
  if moduleName == "" {
    con.Error(c, "模块名称不能为空", "/admin/access/add")
    return
  }

  // 增加数据
  access := models.Access{
    ModuleName: moduleName,
    Type: accessType,
    ActionName: actionName,
    Url: url,
    ModuleId: moduleId,
    Sort: sort,
    Description: description,
    Status: status,
  }
  if models.DB.Create(&access).Error != nil {
    con.Error(c, "增加数据失败", "/admin/access/add")
    return
  }
  con.Success(c, "增加数据成功", "/admin/access")
}

// 编辑权限页面
func (con AccessController) Edit(c *gin.Context) {
  // 1. 获取当前id对应的数据
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "参数错误", "/admin/access")
    return
  }
  access := models.Access{Id: id}
  models.DB.Find(&access)

  // 2. 获取所有顶级模块数据
  accessList := []models.Access{}
  models.DB.Where("module_id", 0).Find(&accessList)

  c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
    "access": access,
    "accessList": accessList,
  })
}

// 编辑权限接口
func (con AccessController) DoEdit(c *gin.Context) {
  // 1. 获取数据
  id, err1 := models.Int(c.PostForm("id"))
  moduleName := strings.Trim(c.PostForm("module_name"), " ")
  actionName := c.PostForm("action_name")
  accessType, err2 := models.Int(c.PostForm("type"))
  url := c.PostForm("url")
  moduleId, err3 := models.Int(c.PostForm("module_id"))
  sort, err4 := models.Int(c.PostForm("sort"))
  status, err5 := models.Int(c.PostForm("status"))
  description := c.PostForm("description")
  if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
    con.Error(c, "传入参数错误", "/admin/access")
    return
  }


  // 2. 先查询后修改
  access := models.Access{Id: id}
  models.DB.Find(&access)
  access.ModuleName = moduleName
  access.Type = accessType
  access.ActionName = actionName
  access.Url = url
  access.ModuleId = moduleId
  access.Sort = sort
  access.Description = description
  access.Status = status
  if models.DB.Save(&access).Error != nil {
    con.Error(c, "修改数据失败", "/admin/access/edit?id=" + models.String(id))
    return
  }
  
  con.Success(c, "修改数据成功", "/admin/access")
}

// 删除权限接口
func (con AccessController) Delete(c *gin.Context) {
  // 1. 获取数据
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/access")
    return
  }

  // 2. 顶级模块还有子数据则不能删除
  access := models.Access{Id: id}
  models.DB.Find(&access)
  if access.ModuleId == 0 { // 表示是顶级模块
    accessList := []models.Access{}
    models.DB.Where("module_id=?", access.Id).Find(&accessList)

    if len(accessList) > 0 {
      con.Error(c, "当前模块下还有菜单或操作, 请先删除", "/admin/access")
      return
    }
  }
  
  // 3. 删除数据
  models.DB.Delete(&access)
  con.Success(c, "删除数据成功", "/admin/access")
}

