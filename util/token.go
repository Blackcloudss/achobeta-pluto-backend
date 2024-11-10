package util

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"tgwp/log/zlog"
	"time"
)

type MyClaims struct {
	Userid string `json:"userid"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

var mySecret = []byte("sb")

type TokenData struct {
	Userid string
	Class  string
	Issuer string
}

func GenToken(data TokenData, t time.Duration) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		data.Userid,
		data.Class,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(t).Unix(), // 过期时间
			Issuer:    data.Issuer,              // 签发人
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
func IdentifyToken(ctx context.Context, Token string) error {
	//解析token
	_, err := util.ParseToken(Token)
	if err != nil {
		zlog.CtxErrorf(ctx, "IdentifyToken err: %v", err)
		return err
	}
	return nil
}
