package config

type Config struct {
	Server   Server     `yaml:"server"`
	Database Database   `yaml:"database"`
	OAuth    OAuth2Cred `yaml:"oauth"`
}
