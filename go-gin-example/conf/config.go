package conf

type ServerConfig struct {
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	ReadTimeOut int    `mapstructure:"read_time_out" json:"read_time_out"`
	WaitTimeOut int    `mapstructure:"wait_time_out" json:"wait_time_out"`

	AppInfo      App            `mapstructure:"app" json:"app"`
	DateBaseInfo DataBaseConfig `mapstructure:"db" json:"db"`
}

type App struct {
	PageSize  int    `mapstructure:"page_size" json:"page_size"`
	JwtSecret string `mapstructure:"jwt_secret" json:"jwt_secret"`
}

type DataBaseConfig struct {
	DBType      string `mapstructure:"db_type" json:"db_type"`
	User        string `mapstructure:"user" json:"user"`
	PassWord    string `mapstructure:"pass_word" json:"pass_word"`
	Host        string `mapstructure:"host" json:"host"`
	Name        string `mapstructure:"name" json:"name"`
	TablePrefix string `mapstructure:"table_prefix" json:"table_prefix"`
}
