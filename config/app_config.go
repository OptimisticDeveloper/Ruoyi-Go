package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"ruoyi-go/utils/R"
)

var Server *server
var Database *database
var Redis *redis
var Jwt *jwt
var XxlJob *xxlJob

type conf struct {
	Svc         server   `yaml:"server"`
	DB          database `yaml:"database"`
	RedisConfig redis    `yaml:"redis"`
	Jwt         jwt      `yaml:"jwt"`
	XxlJob      xxlJob   `yaml:"xxl-job"`
}

type server struct {
	Port     int    `yaml:"port"`
	RunMode  string `yaml:"runMode"`
	LogLevel string `yaml:"logLevel"`
}

type database struct {
	Type            string `yaml:"type"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbName          string `yaml:"dbname"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type jwt struct {
	Secret string `yaml:"secret"`
	JwtTtl int    `yaml:"jwt_ttl"`
}

type xxlJob struct {
	Enabled          bool   `yaml:"enabled"`
	Env              string `yaml:"env"`
	AdminAddress     string `yaml:"admin_address"`
	AccessToken      string `yaml:"access_token"`
	AppName          string `yaml:"app_name"`
	Address          string `yaml:"address"`
	Ip               string `yaml:"ip"`
	Port             int    `yaml:"port"`
	LogPath          string `yaml:"log_path"`
	LogRetentionDays int    `yaml:"log_retention_days"`
	HttpTimeout      int    `yaml:"http_timeout"`
}

func InitAppConfig(dataFile string) {
	filePath := path.Join("", dataFile)
	_, err := os.Stat(filePath)
	if err != nil {
		log.Printf("config file path %s not exist", filePath)
		panic(R.ReturnFailMsg("config file path " + filePath + " not exist"))
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		panic(R.ReturnFailMsg("yamlFile.Get err   " + err.Error()))
	}
	c := new(conf)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		panic(R.ReturnFailMsg("Unmarshal:" + err.Error()))
	}
	log.Printf("load conf success")
	// 绑定到外部可以访问的变量中
	Server = &c.Svc
	Database = &c.DB
	Redis = &c.RedisConfig
	Jwt = &c.Jwt
	XxlJob = &c.XxlJob
}
