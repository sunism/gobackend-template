package config

type Server struct {
	LogConf LogConf `mapstructure:"logconf" json:"logconf" yaml:"logconf"`
}

