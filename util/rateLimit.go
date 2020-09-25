package util

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type LimitInfo struct {
	islocking bool
	lastTimestamp time.Time
}

//限制url的执行速度 最快为1秒一个
type RateLimit struct {
	lock sync.Mutex
	url map[string]*LimitInfo
}

func NewRateLimit()*RateLimit{
	return &RateLimit{
		url:make(map[string]*LimitInfo),
	}
}

func(rl  *RateLimit)getLock(url string)(*LimitInfo,bool){
	rl.lock.Lock()
	defer rl.lock.Unlock()
	if info,ok := rl.url[url];ok{
		if info.islocking{
			return nil,false
		}

		if time.Now().Sub(info.lastTimestamp)< time.Second{
			return nil,false
		}
		info.islocking = true
		return info,true
	}else{
		info := new(LimitInfo)
		rl.url[url]=info
		info.islocking=true
		return info,true
	}

}


func(rl *RateLimit)Limit(next func(c *gin.Context))func(c *gin.Context){
	return func(c *gin.Context) {
		url := c.Request.RequestURI
		info,ok:=rl.getLock(url)
		if !ok{
			return
		}
		defer func() {
			info.lastTimestamp = time.Now()
			info.islocking=false
		}()
		next(c)

	}
}