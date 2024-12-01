// internal/query/config/config.go
package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	DB   db     `yaml:"db"`
}

type db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}
