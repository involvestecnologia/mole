package models

type ReadConfig struct {
	AppName       string        `mapstructure:"app_name"`
	Mongo         Mongo         `mapstructure:"mongo"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
	Notifier      Notifier      `mapstructure:"notifier"`
	Logstash      Logstash      `mapstructure:"logstash"`
}

type Elasticsearch struct {
	Hosts     string `mapstructure:"hosts"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Source    string `mapstructure:"source"`
	BatchSize int    `mapstructure:"batch_size"`
}

type Mongo struct {
	URI     string `mapstructure:"uri"`
	Timeout int    `mapstructure:"timeout"`
}

type Notifier struct {
	URL      string   `mapstructure:"url"`
	Channels []string `mapstructure:"channel"`
}

type Logstash struct {
	URL string `mapstructure:"url"`
}
