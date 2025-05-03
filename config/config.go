package config

type Config struct {
	Host string
	Port int
}

var AppConfig = Config{
	Host: "0.0.0.0",
	Port: 7379,
}
