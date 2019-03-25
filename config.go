package score

import "github.com/spf13/viper"

type Config struct {
	BindAddr string

	DbProvider   string
	DbConnection string
}

func GetConfig() (config Config, err error) {
	v := viper.New()
	v.SetConfigName("score")

	v.SetDefault("BindAddr", "127.0.0.1:6301")
	v.SetDefault("DbProvider", "sqlite3")
	v.SetDefault("DbConnection", "score.db")

	v.AddConfigPath(".")
	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&config)
	return
}
