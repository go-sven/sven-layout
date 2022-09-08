package conf

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"

)

var ProviderSet = wire.NewSet(
	NewMysqlConfig,
	NewRedisConfig,
)

type Config struct {
	WebConf *WebConfig
	MysqlConf *MysqlConfig
	RedisConf *RedisConfig
	LoggerConf *LoggerConfig
}

type WebConfig struct {
	Name string
	Port string
	Mode string
}

type MysqlConfig struct {
	Host string
	Port string
	Username string
	Password string
	Database string
	Parameter string
	MaxIdle int
	MaxOpen int
}

type RedisConfig struct {
	Host string
	Port string
	Password string
	Db int
}

type LoggerConfig struct {
	Filename string
	MaxSize  int
	MaxBackups int
	MaxAge int
}

func InitConf(path string) *Config  {
	viper.SetConfigType("yaml") //设置配置文件格式
	viper.AddConfigPath(path) //设置配置文件的路径
	viper.SetConfigName("config")	//设置配置文件名
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	//返回配置文件的数据
	return  &Config{
		WebConf:   &WebConfig{
			Name: viper.GetString("web.name"),
			Port: viper.GetString("web.port"),
			Mode: viper.GetString("web.mode"),
		},
		MysqlConf : &MysqlConfig{
			Host:       viper.GetString("mysql.host"),
			Port:       viper.GetString("mysql.port"),
			Username:   viper.GetString("mysql.username"),
			Password:   viper.GetString("mysql.password"),
			Database:   viper.GetString("mysql.database"),
			Parameter:  viper.GetString("mysql.parameter"),
			MaxIdle: viper.GetInt("mysql.maxIdle"),
			MaxOpen: viper.GetInt("mysql.maxOpen"),
		},
		RedisConf : &RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			Db:       viper.GetInt("redis.db"),
		},
		LoggerConf: &LoggerConfig{
			Filename:   viper.GetString("logger.filename"),
			MaxSize:    viper.GetInt("logger.maxsize"),
			MaxBackups: viper.GetInt("logger.maxbackups"),
			MaxAge:     viper.GetInt("logger.maxage"),
		},
	}

}

func NewMysqlConfig(config *Config) *MysqlConfig  {
	return config.MysqlConf
}
func NewRedisConfig(config *Config) *RedisConfig  {
	return config.RedisConf
}
