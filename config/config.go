package config

import (
	"fmt"
	"os"

	"github.com/NursiNursi/laundry-apps/utils/common"
	"github.com/NursiNursi/laundry-apps/utils/exceptions"
)

type DbConfig struct {
	Host string
	Port string
	Name string
	User string
	Password string
	Driver string
}

type Config struct {
	DbConfig
}

// Method
func (c *Config) ReadConfig() error {
  err := common.LoadEnv()
	exceptions.CheckErr(err)

	c.DbConfig = DbConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver: os.Getenv("DB_DRIVER"),
	}
	
	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" {
		return fmt.Errorf("missing required env")
	}

	return nil
}

// constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
