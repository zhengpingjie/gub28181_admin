package config

import (
	"admin/http/controller"
	"admin/util"
	"admin/util/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type userRouter struct {
	router *gin.RouterGroup
	webBase_ *webBase
}

func NewRouterUser(r *gin.RouterGroup,c *webBase)*userRouter{
	userRouter_ := &userRouter{}
	userRouter_.webBase_ = c
	userRouter_.router = r
	//初始化user相关路由
	userRouter_.initUserRouter()
	return userRouter_

}

//用户登录信息
type UserLoginReq struct{
	Account string `json:"account"`
	Method string `json:"method"`
	VerifyCode string `json:"verify_code"`
	Pwd string `json:"pwd"`
}

func(u *userRouter)initUserRouter(){
	//路由限速
	speedLimit := util.NewRateLimit()
	//设置路由
	u.router.POST(pathUserLogin,speedLimit.Limit(u.login))

}

func(u *userRouter)login(c *gin.Context){
	var req UserLoginReq
	_=c.ShouldBindJSON(&req)
	verify := util.Rules{
		"Account":{util.NotEmpty()},
		"Method":{util.NotEmpty()},
	}
	if err :=util.Verify(req,verify);err != nil{
		response.OkWithMessage(err.Error(),c)
		return
	}

	switch req.Method {
	case "acount_pwd":
	  		//账号密码登录
	  		if len(strings.Trim(req.Pwd,""))<=0{
	  			response.FailWithMessage("密码不能为空",c)
			}
	        controller.ServerAccountPwdLogin(req.Account,req.Pwd)
	  		break
	case "phone_code":
		//手机号验证码登录
		controller.ServerPhoneCodeLogin(req.Account,req.VerifyCode)
		break
	case "email_code":
		//邮箱验证码登录
		controller.ServerEmailCode(req.Account,req.VerifyCode)
		break
	default:
		response.FailWithDetailed(response.ERROR,req,"",c)
	    return
	}

}