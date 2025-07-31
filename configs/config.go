package configs

import (
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server ServerConfig           `yaml:"server"`
	MySQL  map[string]MySQLConfig `yaml:"mysql"`
	Redis  map[string]RedisConfig `yaml:"redis"`
	ES     ESConfig               `yaml:"elasticsearch"`
	Log    LogConfig              `yaml:"log"`
}

type ServerConfig struct {
	Port string `yaml:"port" default:"8080"`
	Mode string `yaml:"mode" default:"debug"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type ESConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
	Output     string `yaml:"output"`      // 输出位置: stdout, stderr, file
	Dir        string `yaml:"dir"`         // 日志文件目录 (当output为file时)
	FilePrefix string `yaml:"file_prefix"` // 日志文件名前缀(当output为file时)
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	file, err := os.Open("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
