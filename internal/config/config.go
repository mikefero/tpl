/*
Copyright © 2023 Michael Fero

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `yaml:"database" mapstructure:"database"`
	Logger   Logger   `yaml:"logger" mapstructure:"logger"`
}

type Database struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	DBName   string `yaml:"dbname" mapstructure:"dbname"`
}

type Logger struct {
	Level     string `yaml:"level" mapstructure:"level"`
	Filename  string `yaml:"filename" mapstructure:"filename"`
	Retention int    `yaml:"retention" mapstructure:"retention"`
}

func NewConfig() (*Config, error) {
	// Database defaults
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.username", "tpl")
	viper.SetDefault("database.password", "tpl")
	viper.SetDefault("database.dbname", "tpl")

	// Logger defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.filename", "tpl-log.json")
	viper.SetDefault("logger.retention", 60)

	// TPL configuration setup for viper
	viper.SetConfigName("tpl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("TPL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var config Config
	_ = viper.ReadInConfig()
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &config, nil
}

func (c *Config) DataSourceName() string {
	// TODO(fero): Enable TLS
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName)
}
