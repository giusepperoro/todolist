package config

type ServiceConfiguration struct {
	PostgresConnectUrl string `yaml:"postgresConnectUrl"`
	ServerAddressUrl   string `yaml:"serverAddressUrl"`
}
