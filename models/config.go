package models

type Elasticsearch struct {
	Hosts    string `mapstructure:"hosts"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
