package setting

import (
	"github.com/youminghang/go-gin-example/config"
	"gorm.io/gorm"
)

var (
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
)
