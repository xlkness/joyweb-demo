package server

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if global.GVA_CONFIG.System.Env != "develop" {
		//	waitUse, _ := utils.GetClaims(c)
		//	//获取请求的PATH
		//	path := c.Request.URL.Path
		//	obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
		//	// 获取请求方法
		//	act := c.Request.Method
		//	// 获取用户的角色
		//	sub := strconv.Itoa(int(waitUse.AuthorityId))
		//	e := casbinService.Casbin() // 判断策略中是否存在
		//	success, _ := e.Enforce(sub, obj, act)
		//	if !success {
		//		response.FailWithDetailed(gin.H{}, "权限不足", c)
		//		c.Abort()
		//		return
		//	}
		//}
		c.Next()
	}
}

func casbinConfDemo() {
	e, err := casbin.NewEnforcer("path/model.conf", "path/policy.csv")
	e.EnableAutoSave(true)

	sub := "alice"
	obj := "gmtool"
	act := "read"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {

	}

	if ok {
		// pass
	} else {
		// not pass
	}
}
