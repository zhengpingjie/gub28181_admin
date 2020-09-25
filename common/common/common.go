package common

import(
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)
var(
	GVA_VP *viper.Viper
	GVA_LOG *zap.Logger
	GVA_REDIS *redis.Client
	GVA_PG *pg.DB
)