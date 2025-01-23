package global

type System struct {
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	Router_prefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Job_log       string `mapstructure:"job-log" json:"job-log" yaml:"job-log"`
}
