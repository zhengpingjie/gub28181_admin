package config

import (
	"admin/util"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var key = []byte("iioJalStfgGUfU5Xd1wNbUru9ZNVxDjq")
func isAuth(authStr string)bool{
	encryptBytes,err := base64.StdEncoding.DecodeString(authStr)
	if err != nil{
		return false
	}
	plainBytes,err := util.AESDecryptCFB(encryptBytes,key)
	if err != nil{
		return false
	}

	timeStr := string(plainBytes)
	timestamp,err:=strconv.ParseInt(timeStr,10,64)
	if err != nil{
		return false
	}

	return timestamp >= time.Now().Unix() && timestamp<=time.Now().Unix()+15
}

func Auth()gin.HandlerFunc{
	return func(g *gin.Context){
		if isAuth(g.Request.Header.Get("X-Auth")){
			g.Next()
		}else{
			g.Abort()
		}
	}
}
