package initialize

import (
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"github.com/youminghang/go-gin-example/pkg/setting"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("GO_GIN_EXAMPLE_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-debug.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(setting.NacosConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("Nacos配置信息:%v", setting.NacosConfig)
	// viper动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("Nacos配置文件发生变化: %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(setting.NacosConfig)
		zap.S().Infof("配置信息: %v", setting.NacosConfig)
	})
	// 从Nacos中读取配置信息
	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         setting.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	sc := []constant.ServerConfig{
		{
			IpAddr: setting.NacosConfig.Host,
			Port:   uint64(setting.NacosConfig.Port),
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: setting.NacosConfig.DataId,
		Group:  setting.NacosConfig.Group})

	if err != nil {
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &setting.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	fmt.Println(&setting.ServerConfig)
}
