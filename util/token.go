package util

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"tgwp/global"
	"tgwp/log/zlog"
	"time"
)

type MyClaims struct {
	Userid string `json:"userid"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

var mySecret = []byte("AchoBeta")

type TokenData struct {
	Userid string
	Class  string
	Issuer string
	Time   time.Duration
}

func GenToken(data TokenData) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		data.Userid,
		data.Class,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(data.Time).Unix(), // 过期时间
			Issuer:    data.Issuer,                      // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 用于验证令牌是否有效
func IdentifyToken(ctx context.Context, Token string) (TokenData, error) {
	//解析token
	claim, err := ParseToken(Token)
	var data TokenData
	if err != nil {
		zlog.CtxErrorf(ctx, "IdentifyToken err: %v", err)
		return TokenData{}, err
	}
	data.Userid = claim.Userid
	data.Issuer = claim.Issuer
	data.Class = claim.Type
	if claim.Type == "rtoken" {
		data.Time = global.RTOKEN_EFFECTIVE_TIME - time.Duration(time.Now().Unix()-claim.StandardClaims.NotBefore)
	} else {
		data.Time = global.ATOKEN_EFFECTIVE_TIME
	}
	return data, nil
}

func FullToken(class string) (data TokenData) {
	//后期这两个都由雪花算法生成
	data.Issuer = "Zcg"
	data.Userid = "test_id"
	if class == "atoken" {
		data.Time = global.ATOKEN_EFFECTIVE_TIME
		data.Class = "atoken"
	} else {
		data.Time = global.RTOKEN_EFFECTIVE_TIME
		data.Class = "rtoken"
	}
	return
}
