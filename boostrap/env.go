package boostrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	Port      string `mapstructure:"PORT"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
	DBName    string `mapstructure:"DB_NAME"`
	DbSslMode string `mapstructure:"DB_SSLMODE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(".env не найден: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("env не может быть загружен: ", err)
	}

	return &env

}
