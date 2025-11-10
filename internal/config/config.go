package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 数据库和日志配置结构体
 type DBConfig struct {
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Name           string `yaml:"name"`
	Charset        string `yaml:"charset"`
	MaxOpenConns   int    `yaml:"max_open_conns"`
	MaxIdleConns   int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int   `yaml:"conn_max_lifetime"`
}

type LogConfig struct {
	File       string `yaml:"file"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

type Config struct {
	DB  DBConfig  `yaml:"db"`
	Log LogConfig `yaml:"log"`
}

// LoadConfig 加载YAML配置到Config结构体
func LoadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
