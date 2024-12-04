package config

type Config struct {
	Redis Redis `mapstructure:"redis" yaml:"redis" json:"redis"`
}
