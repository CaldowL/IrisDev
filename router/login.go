package router

/*
	扫码登陆
		1.获取一个二维码 开始维护是否被扫码状态
		2.用户扫描二维码 打开扫码页面时,先发送QR被扫描状态
		3.客户端 轮询是否被扫描，返回状态
		4.客户确认后，发送确认状态
		5.轮询查询请求返回确认状态并返回响应数据
*/
import (
	"IrisDev/utils"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
	"time"
)

var userQrHashMap = map[string]string{}

// 0:未扫描 1:已经扫描 2:已经同意 3:拒绝
var userOpenStatusMap = map[string]int{}

func getQrCode() string {
	code := utils.GetRandomString(10)
	userOpenStatusMap[code] = 0
	return code
}

func LoginGetQrCode(ctx iris.Context) {
	_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
		Ret:  0,
		Data: getQrCode(),
	}))
}

var jwtSecretToken = []byte("ly05661265")

type userClaim struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

// generateToken 根据userId生成token
func generateToken(user string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim{
		UserId: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	})
	s, _ := token.SignedString(jwtSecretToken)
	return s
}

// checkToken 检查token是否合法
func checkToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("ly05661265"), nil
	})
	if token.Valid {
		return true, nil
	} else {
		fmt.Println(err.Error())
		return false, err
	}
}

// LoginSetStatus 用户扫码触发
func LoginSetStatus(ctx iris.Context) {
	var qr = ""
	var action = ""
	if ctx.URLParamExists("qr") && ctx.URLParamExists("action") {
		qr = ctx.URLParam("qr")
		action = ctx.URLParam("action")
	} else {
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret:      1,
			ErrorMsg: "invalid qr or action",
		}))
		return
	}
	if action == "open" {
		userOpenStatusMap[qr] = 1
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret: 0,
		}))
		return
	}
	if action == "agree" {
		if !ctx.URLParamExists("user") {
			_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
				Ret:      1,
				ErrorMsg: "agree but invalid user",
			}))
			return
		}
		userOpenStatusMap[qr] = 2
		userQrHashMap[qr] = ctx.URLParam("user")
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret: 0,
		}))
		return
	}
	if action == "refuse" {
		userOpenStatusMap[qr] = 3
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret: 0,
		}))
		return
	}
}

// LoginGetStatus 客户端轮询获取状态接口 0:未扫描 1:已经扫描 2:已经同意
func LoginGetStatus(ctx iris.Context) {
	var qr = ""
	if ctx.URLParamExists("qr") {
		qr = ctx.URLParam("qr")
	} else {
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret:      1,
			ErrorMsg: "invalid user",
		}))
		return
	}
	if userOpenStatusMap[qr] == 2 { // 表示用户已经确认
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret: 0,
			Data: iris.Map{
				"code":  userOpenStatusMap[qr],
				"tip":   "0:未扫描 1:已经扫描 2:已经同意 3:已经拒绝",
				"user":  userQrHashMap[qr],
				"token": generateToken(userQrHashMap[qr]),
			},
		}))
	} else {
		_, _ = ctx.WriteString(utils.JsonFromObj(utils.ResponseBasicBody{
			Ret: 0,
			Data: iris.Map{
				"code": userOpenStatusMap[qr],
				"tip":  "0:未扫描 1:已经扫描 2:已经同意 3:已经拒绝",
			},
		}))
	}

}
