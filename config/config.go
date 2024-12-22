package config

import (
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
	"task-pilot/internal/utils"
)

type ServiceContext struct {
	Config config
	Mysql  *gorm.DB
	Redis  *redis.Client
}

type config struct {
	Name  string `yaml:"Name"`
	Host  string `yaml:"Host"`
	Port  string `yaml:"Port"`
	Mysql struct {
		DataSource string `yaml:"DataSource"`
	}
	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
	}
}

func Init() (*ServiceContext, error) {
	var ctx ServiceContext
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &ctx.Config)
	if err != nil {
		return nil, err
	}
	ctx.Mysql = utils.InitGorm(ctx.Config.Mysql.DataSource)
	ctx.Redis = utils.InitRedis(ctx.Config.Redis.Addr, ctx.Config.Redis.Password, 0)
	return &ctx, nil
}
