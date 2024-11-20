package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *Config

type ServicesConfig struct {
	Id         string `mapstructure:"id"`
	ClientHost string `mapstructure:"clienttHost"`
	ClientPort int    `mapstructure:"clientPort"`
}

type Domain struct {
	Name        string `mapstructure:"name"`
	LoadBalance string `mapstructure:"loadBalance"`
}

type JwtConf struct {
	Secret string `mapstructure:"secret"`
	Exp    int64  `mapstructure:"exp"`
}

type LogConf struct {
	Level string `mapstructure:"level"`
}

type Database struct {
	Url         string `mapstructure:"url"`
	Db          string `mapstructure:"db"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	MinPoolSize string `mapstructure:"minPoolSize"`
	MaxPoolSize string `mapstructure:"maxPoolSize"`
}

type RedisConf struct {
	Addr          string   `mapstructure:"addr"`
	ClusterAddres []string `mapstructure:"clusterAddres"`
	Password      string   `mapstructure:"password"`
	PoolSize      int      `mapstructure:"poolSize"`
	MinIdleConns  int      `mapstructure:"minIdleConns"`
	Host          string   `mapstructure:"host"`
	Port          int      `mapstructure:"port"`
}

type EtcdConf struct {
	Addrs      []string       `mapstructure:"addrs"`
	RWTimeout  string         `mapstructure:"rwTimeout"`
	DiaTimeOut int            `mapstructure:"diaTimeOut"`
	Register   RegisterServer `mapstructure:"register"`
}

type RegisterServer struct {
	Addr    string `mapstructure:"addr"`
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Weight  int    `mapstructure:"weight"`
	Ttl     int64  `mapstructure:"ttl"` // 租约时长
}

type GrocConfgi struct {
	Addr string `mapstructure:"addr"`
}

type Config struct {
	Log        LogConf                   `mapstructure:"log"`
	Port       int                       `mapstructure:"port"`
	WsPort     int                       `mapstructure:"WsPort"`
	MetricPort int                       `mapstructure:"metricPort"` // 实时监控端口
	HttpPort   int                       `mapstructure:"httpPort"`
	AppName    string                    `mapstructure:"appName"`
	Database   Database                  `mapstructure:"database"`
	Jwt        JwtConf                   `mapstructure:"jwt"`
	Grpc       GrocConfgi                `mapstructure:"grpc"`
	Etcd       EtcdConf                  `mapstructure:"etcd"`
	Domain     map[string]Domain         `mapstructure:"domain"`
	Services   map[string]ServicesConfig `mapstructure:"services"`
}

// 加载配置
func IninConfig(configFile string) {
	Conf = new(Config)
	v := viper.New()
	v.SetConfigFile(configFile)

	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {

		// 文件变化后重新加载配置 输出日志
		err := v.Unmarshal(Conf)
		log.Println("配置文件发生改变")
		if err != nil {
			panic(fmt.Errorf("配置文件变化,解析配置文件出错,Unmarshal err", err))
		}
		log.Println("配置重载成功")
	}) 

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件出错,err", err))
	}

	// 解析配置文件
	err = v.Unmarshal(Conf)
	if err != nil {
		panic(fmt.Errorf("解析配置文件出错,Unmarshal err", err))
	}
}
