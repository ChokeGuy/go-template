package config

import (
	"github.com/spf13/viper"
)

// Struct config
type Config struct {
	SERVER_ADDRESS          string `mapstruture:"SERVER_ADDRESS"`
	PORT                    string `mapstruture:"PORT"`
	KAFKA_BROKER            string `mapstruture:"KAFKA_URL"`
	KAFKA_SECURITY_PROTOCOL string `mapstruture:"KAFKA_SECURITY_PROTOCOL"`
	KAFKA_SASL_MECHANISM    string `mapstruture:"KAFKA_SASL_MECHANISM"`
	AWS_ACCESS_KEY_ID       string `mapstruture:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY   string `mapstruture:"AWS_SECRET_ACCESS_KEY"`
	AWS_REGION              string `mapstruture:"AWS_REGION"`
	WRITER_POSTGRES_URL     string `mapstruture:"WRITER_POSTGRES_URL"`
	READER_POSTGRES_URL     string `mapstruture:"READER_POSTGRES_URL"`
	INIT_TOPICS             bool   `mapstruture:"INIT_TOPICS"`
	KAFKA_GROUP_ID          string `mapstruture:"KAFKA_GROUP_ID"`
	PREFIX_PATH             string `mapstruture:"PREFIX_PATH"`
}

// Use viper to load config from file .env
func InitConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	config := &Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
