package config

type Config struct {
	Server Server `yaml:"server"`
	OAuth  OAuth2Cred
}
