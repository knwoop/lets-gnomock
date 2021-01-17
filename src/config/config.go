package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	MysqlHost     string `envconfig:"MYSQL_HOST" default:"127.0.0.1"`
	MysqlPort     string `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlUser     string `envconfig:"MYSQL_USER" default:"root"`
	MysqlPass     string `envconfig:"MYSQL_PASS" default:""`
	MysqlDatabase string `envconfig:"MYSQL_DATABASE" default:""`
}

func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, fmt.Errorf("failed to process envconfig: %w", err)
	}

	return &env, nil
}
