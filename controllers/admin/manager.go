package admin

import (
	"net/http"
	"strings"
	"xiaomishop/models"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
  BaseController
}

// 管理员界面
func (con ManagerController) Index(c *gin.Context) {
  managerList := []models.Manager{}
  models.DB.Preload("Role").Find(&managerList)

  c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
    "managerList": managerList,
  })
}

// 增加管理员页面
func (con ManagerController) Add(c *gin.Context) {
  // 获取所有角色
  roleList := []models.Role{}
  models.DB.Find(&roleList)
  
  c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
    "roleList": roleList,
  })
}

// 增加管理员接口
func (con ManagerController) DoAdd(c *gin.Context) {
  // 1. 获取角色id
  roleId, err := models.Int(c.PostForm("role_id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/manager/add")
    return
  }

  // 2. 获取其他数据
  username := strings.Trim(c.PostForm("username"), " ")
  password := strings.Trim(c.PostForm("password"), " ")
  email := strings.Trim(c.PostForm("email"), " ")
  mobile := strings.Trim(c.PostForm("mobile"), " ")
  if len(username) < 2 || len(password) < 6 {
    con.Error(c, "用户名或密码长度不合法", "/admin/manager/add")
    return
  }

  // 3. 判断管理员是否存在 (用户名是否存在)
  managerList := []models.Manager{}
  models.DB.Where("username = ?", username).Find(&managerList)
  if len(managerList) > 0 {
    con.Error(c, "此管理员已经存在", "/admin/manager/add")
    return
  }

  // 4. 插入数据库
  manager := models.Manager{
    Username: username,
    Password: models.Md5(password), // 密码进行md5加密
    Email: email,
    Mobile: mobile,
    RoleId: roleId,
    Status: 1,
    AddTime: int(models.GetUnix()),
  }
  if models.DB.Create(&manager).Error != nil {
    con.Error(c, "增加管理员失败", "/admin/manager/add")
    return
  }

  con.Success(c, "增加管理员成功", "/admin/manager")
}

// 修改管理员页面
func (con ManagerController) Edit(c *gin.Context) {
  // 1. 获取数据
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/manager")
    return
  }

  // 2. 获取当前id对应的管理员数据
  manager := models.Manager{Id: id}
  models.DB.Find(&manager)

  // 3. 获取所有的角色数据
  roleList := []models.Role{}
  models.DB.Find(&roleList)

  c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
    "manager": manager,
    "roleList": roleList,
  })
}

// 修改管理员接口
func (con ManagerController) DoEdit(c *gin.Context) {
  // 1. 获取管理员id和角色id
  managerId, err1 := models.Int(c.PostForm("manager_id"))
  if err1 != nil {
    con.Error(c, "传入数据错误", "/admin/manager")
    return
  }
  roleId, err2 := models.Int(c.PostForm("role_id"))
  if err2 != nil {
    con.Error(c, "传入数据错误", "/admin/manager")
    return
  }

  // 2. 获取其他数据
  password := strings.Trim(c.PostForm("password"), " ")
  email := strings.Trim(c.PostForm("email"), " ")
  mobile := strings.Trim(c.PostForm("mobile"), " ")
  if len(mobile) > 11 {
    con.Error(c, "mobile长度不合法", "/admin/manager/edit?id=" + models.String(managerId))
    return
  }
  
  // 3. 先获取数据再修改
  manager := models.Manager{Id: managerId}
  models.DB.Find(&manager)
  manager.Email = email
  manager.Mobile = mobile
  manager.RoleId = roleId
  // 注意: 判断密码是否空, 为空表示不修改密码, 不为空表示修改密码
  if password != "" {
    // 判断密码长度是否合法
    if len(password) < 6 {
      con.Error(c, "密码长度不合法", "/admin/manager/edit?id=" + models.String(managerId))
      return
    }
    manager.Password = models.Md5(password)
  }
  // 修改管理员
  if models.DB.Save(&manager).Error != nil {  
    con.Error(c, "修改数据失败", "/admin/manager/edit?id=" + models.String(managerId))
    return
  }

  con.Success(c, "修改数据成功", "/admin/manager")
}

// 删除管理员接口
func (con ManagerController) Delete(c *gin.Context) {
  id, err := models.Int(c.Query("id"))
  if err != nil {
    con.Error(c, "传入数据错误", "/admin/manage")
    return
  } else {
    manager := models.Manager{Id: id}
    models.DB.Delete(&manager)
    con.Success(c, "删除数据成功", "/admin/manage")
  }
}