package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 全局变量, 用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init(filename string) (err error) {
	viper.SetConfigFile(filename)
	// viper.SetConfigFile("./config.yaml")
	// viper.SetConfigName("config")               // 指定配置文件 不需要带后缀
	// viper.SetConfigType("yaml")                 //指定配置类型（专用于从远程获取配置信息时指定的配置）
	// viper.AddConfigPath(".")                    //指定查找配置文件的路径
	if err = viper.ReadInConfig(); err != nil { // 读取配置信息
		zap.L().Error("读取配置文件错误", zap.Error(err))
		return
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("读取信息反序列化错误", zap.Error(err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("读取信息反序列化错误", zap.Error(err))
		}
		zap.L().Info("配置文件发生了修改")
		fmt.Println("配置文件发生了修改")
	})
	return
}
