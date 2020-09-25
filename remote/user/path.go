package user

const(
	//添加用户
	PathUserReg = "/servers/user/register"
	//用户登录
	PathUserLogin = "/servers/user/login"
	//获取登录验证码
	PathUserGetVerifyCode = "/servers/user/verify"
	//退出登录
	PathUserLoginOut = "/servers/user/loginOut"
	//获取用户信息
	PathUserGetInfo = "servers/user/getUserInfo"
	//设置登录密码
	PathUserSetPwd = "servers/user/setUserPwd"
	//修改登录密码
	PathUserChangePwd = "servers/user/changePwd"

)

type LoginUserInfo struct{
	UserName string `json:"username"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
	IsEnable int8 `json:"is_enable"`
}

type LoginRequest struct {
	Account string `json:"account`
	Pwd string `json:"pwd"`
	Phone string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}