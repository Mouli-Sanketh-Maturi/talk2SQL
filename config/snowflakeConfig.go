package config

type SnowflakeConfig struct {
	Account   string `json:"account"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Warehouse string `json:"warehouse"`
}
