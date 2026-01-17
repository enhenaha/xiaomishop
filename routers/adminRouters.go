package routers

import (
	"xiaomishop/controllers/admin"
	middlewars "xiaomishop/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
  adminRouters := r.Group("/admin", middlewars.InitAdminAuthMiddleware)
  {
    // 后台首页
    adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

    // 登录页面
    adminRouters.GET("/login", admin.LoginController{}.Index)
    adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
    adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
    adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)

    // 管理员页面
    adminRouters.GET("/manager", admin.ManagerController{}.Index)
    adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
    adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
    adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
    adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
    adminRouters.GET("/manager/delete", admin.RoleController{}.Delete)

    // 轮播图页面
    adminRouters.GET("/focus", admin.FocusController{}.Index)
    adminRouters.GET("/focus/add", admin.FocusController{}.Add)
    adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)

    // 角色页面
    adminRouters.GET("/role", admin.RoleController{}.Index)
    adminRouters.GET("/role/add", admin.RoleController{}.Add)
    adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
    adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
    adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)
    adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
    adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
    adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)

    // 权限
    adminRouters.GET("/access", admin.AccessController{}.Index)
    adminRouters.GET("/access/add", admin.AccessController{}.Add)
    adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)
    adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
    adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
    adminRouters.GET("/access/delete", admin.AccessController{}.Delete)
  }
}