package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func filtterWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("receive request[%v]\n", c.Request.URL.Path)
		c.Next()
	}
}

func jwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		c.Next()
		return
		if c.Request.URL.Path != "/login" && token == "" {
			fmt.Printf("path(%v) not found token:%v\n", c.Request.URL.Path, c.Request.Header)
			c.JSON(300, map[string]interface{}{
				"code":    300,
				"message": "token is empty",
				"payload": nil,
			})
			c.Abort()
			return
		}

		// 校验token是否白名单

		// todo：jwt校验token合法

		// todo：jwt解析token中的uuid

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
