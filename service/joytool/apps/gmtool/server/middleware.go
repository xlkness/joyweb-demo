package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"joytool/apps/gmtool/model"
	"joytool/lib/token"
	"time"
)

func filtterWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("receive request[%v]\n", c.Request.URL.Path)
		c.Next()
	}
}

func permission() func(ctx *model.MyContext) {
	return func(ctx *model.MyContext) {

	}
}

func jwtMiddleWare() func(ctx *model.MyContext) {
	return func(ctx *model.MyContext) {
		c := ctx.GetGinContext()
		tokenStr := c.Request.Header.Get("x-token")
		//c.Next()
		//return
		if tokenStr == "" {
			fmt.Printf("path(%v) not found token:%v\n", c.Request.URL.Path, c.Request.Header)
			c.JSON(300, map[string]interface{}{
				"code":    300,
				"message": "token is empty",
				"payload": nil,
			})
			c.Abort()
			return
		}

		checkedToken, err := token.ValidToken(tokenStr)
		if err != nil {
			fmt.Printf("valid token(%v) not pass, error:%v\n", tokenStr, err)
			c.JSON(300, map[string]interface{}{
				"code":    300,
				"message": err.Error(),
				"payload": nil,
			})
			c.Abort()
			return
		}

		checkedClaims := checkedToken.Claims.(*token.RegisteredTokenClaims)
		fmt.Printf("valid token(%v) pass, user:%v,  expire:%v, issuer:%v\n",
			tokenStr, checkedClaims.User, checkedClaims.ExpiresAt, checkedClaims.Issuer)

		// 校验token是否白名单

		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Content-type", "Access-Token", "Authorization"},
		MaxAge:          6 * time.Hour,
	}
	return cors.New(c)
}
