package redistool

import (
	"admin/common/common"
	"github.com/go-redis/redis"
	"os"
)

//
func Redis() {
	redisConfig :=  common.GVA_VP.GetStringMap("redis")
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig["addr"].(string),
		Password: redisConfig["password"].(string), // no password set
		DB:       redisConfig["db"].(int),       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		common.GVA_LOG.Error(err.Error())
		os.Exit(0)
	} else {
		common.GVA_LOG.Info("redis connect ping success")
		common.GVA_REDIS = client
	}
}