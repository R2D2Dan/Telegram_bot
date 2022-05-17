package config

import "github.com/spf13/viper"

type Messages struct {
	Start   string `mapstructure:"start"`
	Unknown string `mapstructure:"unknown"`
}

func InitConfig() (*Messages, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var m Messages

	if err := getMessage(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func getMessage(m *Messages) error {
	if err := viper.UnmarshalKey("messages", &m); err != nil {
		return err
	}
	return nil
}

func setUpViper() error {
	viper.AddConfigPath("Data")
	viper.SetConfigName("messages")

	return viper.ReadInConfig()
}
