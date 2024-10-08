package global

type System struct {
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	Router_prefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
