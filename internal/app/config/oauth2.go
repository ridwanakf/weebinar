package config

type OAuth2Cred struct {
	ClientID     string `env:"OAUTH_CLIENT_ID"`
	ClientSecret string `env:"OAUTH_CLIENT_SECRET"`
	RandomState  string `env:"OAUTH_RANDOM_STATE"`
}
