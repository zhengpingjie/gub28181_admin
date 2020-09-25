package config

import(
	"admin/common/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
)

type webBase struct{
	router *gin.RouterGroup
	lock sync.Mutex
}


func (c *webBase)initBaseCmd(addr string){
	//白名单


	err:=c.init()
	if err != nil {
		common.GVA_LOG.Fatal("web client init fatal",zap.Any("err",err))
	}
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//验证header
	router.Use(Auth())
	//
}