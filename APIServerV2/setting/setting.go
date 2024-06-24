package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init() error {
	// 1.设置yaml文件路径
	viper.SetConfigFile("./config/config.yaml")

	//2.实时监控yaml文件是否被修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Print("文件被修改了")
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(fmt.Errorf("%v", err))
		}
	})

	//3.读取yaml文件参数
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		return err
	}
	return nil
}
