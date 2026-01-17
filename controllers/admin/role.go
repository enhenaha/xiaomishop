package admin

import (
	"net/http"
	"strings"
	"xiaomishop/models"

	"github.com/gin-gonic/gin"
)

type RoleController struct{
  BaseController
}

// 角色列表页面
func (con RoleController) Index(c *gin.Context) {
  // 查询数据
  roleList := []models.Role{}
  models.DB.Find(&roleList)

  c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
    "roleList": roleList,
  })
}


// 增加角色页面
func (con RoleController) Add(c *gin.Context) {
  c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

// 增加角色接口
func (con RoleController) DoAdd(c *gin.Context) {
  // 1. 获取数据并判断
  title := strings.Trim(c.PostForm("title"), " ")
  description := strings.Trim(c.PostForm("description"), " ")
  if title == "" {
    con.Error(c, "角色的标题不能为空", "/admin/role/add")
    return
  }

  // 2. 插入数据
  role := models.Role{
    Title: title,
    Description: description,
    Status: 1,
    AddTime: int(models.GetUnix()),
  }
  err := models.DB.Create(&role).Error

  // 3. 判断插入结果
  if err != nil {
    con.Error(c, "增加角色失败", "/admin/role/add")
  } else {
    con.Success(c, "增加角色成功", "/admin/role")
  }
}

// 编辑角色页面
func (con RoleController) Edit(c *gin.Context) {
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/role")
    return
  } else {
    role := models.Role{Id: id}
    models.DB.Find(&role)

    c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
      "role": role,
    })
  }
}

// 修改角色接口
func (con RoleController) DoEdit(c *gin.Context) {
  // 1. 获取数据并判断
  id, err := models.Int(c.PostForm("id")) // 从隐藏表单域中获取数据
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/role")
    return
  }
  title := strings.Trim(c.PostForm("title"), " ")
  description := strings.Trim(c.PostForm("description"), " ")
  if title == "" {
    con.Error(c, "角色的标题不能为空", "/admin/role/edit")
    return
  }

  // 2. 查询要修改的数据并修改
  role := models.Role{Id: id}
  models.DB.Find(&role)

  role.Title = title
  role.Description = description
  // 完整写法:
  // err := models.DB.Save(&role).Error
  // if err != nil {}
  // 简写:
  if models.DB.Save(&role).Error != nil {
    con.Error(c, "修改数据失败", "/admin/role/edit?id=" + models.String(id))
    return
  } else {
    con.Success(c, "修改数据成功", "/admin/role")
    return
  }
}

// 删除角色接口
func (con RoleController) Delete(c *gin.Context) {
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/role")
    return
  } else {
    role := models.Role{Id: id}
    models.DB.Delete(&role)
    con.Success(c, "删除数据成功", "/admin/role")
    return
  }
}


// 角色授权页面
func (con RoleController) Auth(c *gin.Context) {
  // 1. 获取角色id
  roleId, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/role")
    return
  }

  // 2. 获取所有的权限
  accessList := []models.Access{}
  models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)

  // 3. 获取当前角色拥有的权限, 并将所有权限id存入map对象中
  roleAccess := []models.RoleAccess{}
  models.DB.Where("role_id=?", roleId).Find(&roleAccess)

  roleAccessMap := make(map[int]int)
  for _, v := range roleAccess {
    roleAccessMap[v.AccessId] = v.AccessId
  }

  // 4. 遍历所有的权限, 判断当前权限id是否在角色权限的map对象中, 如果存在则给当前权限增加checked属性。
  // 通过for...range无法修改切片中的结构体元素数据, 需要用for循环。
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

  c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
    "roleId": roleId,
    "accessList": accessList,
  })
}

// 角色授权接口
func (con RoleController) DoAuth(c *gin.Context) {
  // 1. 获取角色id
  roleId, err := models.Int(c.PostForm("role_id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/role")
    return
  }

  // 2. 获取所有勾选的权限id
  // c.PostForm()返回的是字符串, c.PostFormArray()返回的是字符串切片。
  accessIds := c.PostFormArray("access_node[]")

  // 3. 删除当前角色对应的所有权限
  roleAccess := models.RoleAccess{}
  models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

  // 4. 给当前角色插入新的所有权限 (插入role_access表)
  for _, v := range accessIds {
    roleAccess.RoleId = roleId
    accessId, _ := models.Int(v)
    roleAccess.AccessId = accessId
    models.DB.Create(&roleAccess)
  }

  con.Success(c, "授权成功", "/admin/role")
}