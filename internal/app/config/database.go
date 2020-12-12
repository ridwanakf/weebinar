package config

type Database struct {
	Address      string `yaml:"address" env:"DATABASE_URL"`
	Driver       string `yaml:"driver" env:"DATABASE_DRIVER"`
	MaxConns     int    `yaml:"max_conns" env:"DATABASE_MAX_CONNS"`
	MaxIdleConns int    `yaml:"max_idle_conns" env:"DATABASE_MAX_IDLE"`
}
