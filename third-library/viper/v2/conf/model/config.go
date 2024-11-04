package model

// 解析配置文件汇总
type Configurations struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

// 解析应用基本配置
type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

// 解析MySQL配置
type MySQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	DBName   string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
}
