package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	// "testing"
	"gopkg.in/yaml.v2"
)

// Config 配置对象
var Config Conf

// Conf 配置结构体
type Conf struct {
	Env         string `yaml:"env"`
	AppName     string `yaml:"app_name"`
	Mode        string `yaml:"mode"`
	DbLogEnable bool   `yaml:"db_log_enable"`
	GinHost     string `yaml:"gin_host"`

	Db map[string]struct {
		Driver      string        `yaml:"driver"`
		Uri         string        `yaml:"uri"`
		MaxOpenConn int           `yaml:"max_open_conn"`
		MaxIdleConn int           `yaml:"max_idle_conn"`
		MaxLifetime time.Duration `yaml:"max_life_time"`
	} `yaml:"db"`
	Redis map[string]struct {
		Uri       string
		Auth      string
		Db        int
		MaxIdle   int
		MaxActive int
	}
}

var gConfigName string

func init() {
	testing.Init()
	fmt.Println("config init")
	flag.StringVar(&gConfigName, "conf", "./config/conf.yaml", "config name")
	flag.Parse()

	fmt.Println("config name ", gConfigName)
	ParseYaml(gConfigName, &Config)
}

// ParseYaml 解析yaml配置文件
func ParseYaml(file string, configRaw interface{}) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}

	err = yaml.Unmarshal(content, configRaw)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
}
