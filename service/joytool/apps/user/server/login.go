package server

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"joytool/apps/user/model"
	"log"
	"time"
)

type AuthorizeData struct {
}

type TokenData struct {
}

type LoginData struct {
	UserName string   `json:"username"`
	PassWord string   `json:"password"`
	Modules  []string `json:"modules"`
}

// gopkg.in/oauth2.v3
// github.com/go-oauth2/oauth2

var srv = oauth()

func Authorize(ctx *model.MyContext) {
	err := srv.HandleAuthorizeRequest(ctx.GetGinContext().Writer, ctx.GetGinContext().Request)
	if err != nil {
		fmt.Printf("HandleAuthorizeRequest error:%v\n", err)
		ctx.RespFailMessage(300, err.Error())
	}
}

func Token(ctx *model.MyContext) {
	srv.HandleTokenRequest(ctx.GetGinContext().Writer, ctx.GetGinContext().Request)
}

var (
	tokenSign = []byte("token_sign")
)

func validToken(token string) (*jwt.Token, error) {
	checkToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSign, nil
	})
	if err != nil {
		return nil, fmt.Errorf("ParseWithClaims token (%v) error:%v", token, err)
	} else {
		if !checkToken.Valid {
			return nil, fmt.Errorf("ParseWithClaims token (%v) invalid", token)
		}
		//else {
		//	claims := checkToken.Claims.(*jwt.RegisteredClaims)
		//	if claims.ExpiresAt.Before(time.Now()) {
		//		return nil, fmt.Errorf("token expired:%v", claims.ExpiresAt)
		//	}
		//}
	}

	return checkToken, nil
}

func getToken(expire time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(tokenSign)
	if err != nil {
		return "", fmt.Errorf("token SignedString error:%v", err)
	}

	return ss, nil
}

func Login(ctx *model.MyContext, loginData *LoginData) {
	log.Printf("receive login data:%v,%v,%+v", loginData.UserName, loginData.PassWord, loginData.Modules)

	// todo：校验用户名密码

	token, err := getToken(time.Hour * 24 * 30)
	if err != nil {
		ctx.RespFailMessage(300, err.Error())
		return
	}

	checkedToken, err := validToken(token)
	if err != nil {
		fmt.Printf("valid token(%v) not pass, error:%v\n", token, err)
	} else {
		checkedClaims := checkedToken.Claims.(*jwt.RegisteredClaims)
		fmt.Printf("valid token(%v) pass, expire:%v, issuer:%v\n", token, checkedClaims.ExpiresAt, checkedClaims.Issuer)
	}

	// 获取模块权限
	modulesPermission := make([]interface{}, 0, len(loginData.Modules))
	for _, v := range loginData.Modules {
		modulesPermission = append(modulesPermission, map[string]interface{}{
			"name":       v,
			"permission": "read",
		})
	}

	resp := map[string]interface{}{
		"username": loginData.UserName,
		"token":    token,
		"modules":  modulesPermission,
	}
	ctx.RespSuccessJson(resp)
}
