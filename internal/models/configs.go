package models

type Config struct {
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
	RedisParams    RedisParams    `json:"redis_params"`
}

type AppParams struct {
	GinMode    string `json:"Gin_mode"`
	PortRun    string `json:"Port_run"`
	ServerUrl  string `json:"Server_url"`
	ServerName string `json:"Server_name"`
}

type PostgresParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
}

type RedisParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
