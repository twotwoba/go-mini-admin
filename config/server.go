package config

type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // dev/beta/prod
}
