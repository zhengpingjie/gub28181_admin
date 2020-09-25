package main

import (
	"admin/common/common"
	"admin/common/pgsqltool"
	"admin/common/redistool"
	"admin/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func main(){
//初始化配置文件
  initYmal()
  //初始化日志
	initLog()
  //初始化数据库连接
	initDb()


}

func initYmal(){
	//获取项目的执行路径
	path,err:= os.Getwd()
	if err != nil{
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path)
	//设置配置文件路径
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	//读取配置文件
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}else{
		common.GVA_VP = config
	}
}

//初始化日志
func initLog(){
	debug := util.Opt.Debug
	output_path := common.GVA_VP.GetString("app.logPath")
	if debug {
		logger := util.NewZapDevelopmenet(output_path)
		zap.ReplaceGlobals(logger)
		common.GVA_LOG = logger
	} else {
		logger := util.NewZapProduction(output_path)
		zap.ReplaceGlobals(logger)
		common.GVA_LOG = logger
	}

}

//初始化DB
func initDb(){
	//初始化redis
	redistool.Redis()
	//初始化pgsql'
	pgsqltool.Pgsql()
}