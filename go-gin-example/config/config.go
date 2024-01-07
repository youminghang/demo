package config

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataid" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}
type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
type ServerConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	Host         string       `mapstructure:"host" json:"host"`
	Tags         []string     `mapstructure:"tags" json:"tags"`
	MysqlInfo    MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	ConsulInfo   ConsulConfig `mapstructure:"consul" json:"consul"`
	JWTInfo      JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ReadTimeOut  string       `mapstructure:"read_time_out" json:"read_time_out"`
	WriteTimeOut string       `mapstructure:"write_time_out" json:"write_time_out"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
