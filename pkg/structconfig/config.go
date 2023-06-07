package structconfig

import (
	"fmt"
)

// Env contains all environment variable
type Env struct {
	PostgresLog sqlConf `mapstructure:"POSTGRES_LOG_DB"`
	DiscordApp  App3rd  `mapstructure:"DISCORD_APP"`
	FacebookApp App3rd  `mapstructure:"FACEBOOK_APP"`
}

type sqlConf struct {
	Addr       string `mapstructure:"ADDR"`
	DBNet      string `mapstructure:"DB_NET"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

type App3rd struct {
	Key    string `mapstructure:"KEY"`
	Secret string `mapstructure:"SECRET"`
}

func (conf sqlConf) FormatDSN() string {
	postgresDsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		conf.DBUser,
		conf.DBPassword,
		conf.Addr,
		conf.DBName)

	return postgresDsn
}
