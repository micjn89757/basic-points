package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var logger *zap.Logger
var err error

var mySigningKey = []byte("base")

type MyCustomClaims struct {
	Foo string `json:"foo"` // 自定义字段
	jwt.RegisteredClaims  // jwt标准字段
}

// 根据自定义claims生成token
func SetCustomToken() (string, error){
	sugar := logger.Sugar()
	defer sugar.Sync()
	// create claims with mutiple filed polulated
	claims := &MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{ 
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt: jwt.NewNumericDate(time.Now()),  // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			Issuer: "test", // 签发人
			Subject: "somebody", // 主题
			ID: "1", // 编号
			Audience: []string{"somebody_else"}, // 受众
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 生成jwt token
	jwt, err := token.SignedString(mySigningKey) // 进行签名
	if err != nil {
		sugar.Errorf("signed jwt token failed: %s", err.Error())
		return jwt, err
	}

	return jwt, nil
}


// 解析自定义的claim
func ValidCustomToken(tokenString string) bool {
	sugar := logger.Sugar()
	defer sugar.Sync()
	// jwt.WithLeeway 如果一个jwt过期了但是过期不久，设置一个余地，在过期后的5s内还是可以处理
	// 例如设置的过期时间是30s, 但是你可能有时候在30s后才会处理它
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token)(i any, err error) {
		return mySigningKey, nil
	}, jwt.WithLeeway(5*time.Second))  

	if err != nil {
		sugar.Errorf("parse jwt err: %s", err.Error())
		return false
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		sugar.Infoln(claims.Foo, claims.RegisteredClaims.Issuer)
		return true
	} else {
		sugar.Error("unknown claims type, cannot proceed")
		return false 
	}
}

// 使用jwt默认claims生成Token
func SetToken() (string, error) {
	sugar := logger.Sugar()
	defer sugar.Sync()

	claims := &jwt.RegisteredClaims{
		Issuer: "test",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtStr, err := token.SignedString(mySigningKey)
	if err != nil {
		sugar.Errorf("%s", err.Error())
		return jwtStr, err
	}

	return jwtStr, nil
}


// 解析标准claim生成的jwt
func ValidToken(jwtStr string) bool {
	sugar := logger.Sugar()
	defer sugar.Sync()

	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (any, error) {
		return mySigningKey, nil
	})

	if err != nil {
		sugar.Errorf("failed %s", err.Error())
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sugar.Info(claims.GetIssuer())
		return true
	}else {
		return false
	}

}

func init() {
	logger, err = zap.NewProduction()

	if err != nil {
		panic(err)
	}
}