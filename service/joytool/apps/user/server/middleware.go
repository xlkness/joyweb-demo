package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"joytool/apps/user/model"
	"joytool/lib/token"
	"time"
)

func jwtMiddleWare() func(ctx *model.MyContext) {
	return func(ctx *model.MyContext) {
		c := ctx.GetGinContext()

		if c.Request.URL.Path == "/login" {
			return
		}

		tokenStr := c.Request.Header.Get("x-token")
		//c.Next()
		//return
		if tokenStr == "" {
			fmt.Printf("path(%v) not found token:%v\n", c.Request.URL.Path, c.Request.Header)
			ctx.RespFailMessage(300, "token is empty")
			c.Abort()
			return
		}

		checkedToken, err := token.ValidToken(tokenStr)
		if err != nil {
			fmt.Printf("valid token(%v) not pass, error:%v\n", tokenStr, err)
			ctx.RespFailMessage(300, "token is invalid")
			c.Abort()
			return
		}

		checkedClaims := checkedToken.Claims.(*token.RegisteredTokenClaims)
		fmt.Printf("valid token(%v) pass, user:%v,  expire:%v, issuer:%v\n",
			tokenStr, checkedClaims.User, checkedClaims.ExpiresAt, checkedClaims.Issuer)

		// 校验token是否白名单
		c.Set("claims", checkedClaims)
		//c.Request.WithContext(context.WithValue(context.Background(), "claims", checkedClaims))
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
