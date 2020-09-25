package controller

import (
	"admin/common/common"
	"admin/common/servers"
	"admin/remote/user"
)

//账号密码登录
func ServerAccountPwdLogin(account,pwd string)(*user.LoginUserInfo,error){
	common.GVA_LOG.Debug("ctrl:ServerAccountPwdLogin")
	request := &user.LoginRequest{
		Account:account,
		Pwd:pwd,
	}
	servers.SerResquest(user.PathUserLogin,request,true)


}

//手机号验证码登录
func ServerPhoneCodeLogin(phone,code string){

}
//邮箱验证码登录
func ServerEmailCode(email,code string){

}