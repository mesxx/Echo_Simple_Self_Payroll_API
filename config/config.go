package config

import (
	"os"
	"strconv"

	"github.com/mesxx/Echo_Simple_Self_Payroll_API/config/postgres"

	"gorm.io/gorm"
)

type (
	config struct {
	}

	Config interface {
		ServiceName() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)

	return port
}

func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
