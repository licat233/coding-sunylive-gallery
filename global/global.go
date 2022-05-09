package global

import (
	"sunylive-gallery/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	Log         *zap.SugaredLogger
	DB          *gorm.DB
	Redis       *redis.Client
	RootPath    string
	ReleaseMode bool
)
