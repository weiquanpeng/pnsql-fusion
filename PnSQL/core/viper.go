package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xiaohongshu/PnSql/server/global"
)

func Viper(path ...string) {
	var config = "./config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if v.Unmarshal(global.P_cfg) != nil {
		fmt.Printf("无法解析配置到结构体: %s", err)
	}
}
