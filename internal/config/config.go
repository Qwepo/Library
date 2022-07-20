package config

import (
	"io/fs"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Mysql  MsqlConfig   `yaml:"db"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type MsqlConfig struct {
	DBname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *Config) SetDefault() {
	c.Server.Port = "8080"
	c.Mysql.DBname = os.Getenv("MSQL_DBNAME")
	c.Mysql.Username = os.Getenv("MSQL_USERNAME")
	c.Mysql.Password = os.Getenv("MSQL_PASSWORD")

}

func LoadConfig() Config {
	var conf Config
	conf.SetDefault()
	data, err := ioutil.ReadFile("config.yaml")
	if os.IsNotExist(err) {

		if data, err := yaml.Marshal(conf); err == nil {
			_ = ioutil.WriteFile("config.yaml", data, fs.ModePerm)
		}
	} else {
		_ = yaml.Unmarshal(data, &conf)
	}
	return conf
}
