package config

type OAuth2Cred struct {
	ClientID     string `yaml:"client_id" env:"OAUTH_CLIENT_ID"`
	ClientSecret string `yaml:"client_secret" env:"OAUTH_CLIENT_SECRET"`
	RandomState  string `yaml:"random_state" env:"OAUTH_RANDOM_STATE"`
}
