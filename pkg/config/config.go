package config

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	RTPNumber     float64
}

func NewConfig(path string, configType, configName string, RptNumber float64) (Config, error) {
	var config Config
	if err := config.InitConfig(path, configType, configName); err != nil {
		return config, err
	}
	config.RTPNumber = RptNumber
	return config, nil
}
func (c *Config) InitConfig(path, configType, configName string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}
	return nil

}
