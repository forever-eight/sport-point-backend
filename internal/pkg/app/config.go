package app

import "github.com/spf13/viper"

func (a *App) initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
