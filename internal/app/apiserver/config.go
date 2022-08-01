package apiserver

//Config
type Config struct {
	BindAddress string `toml:"bind_address"`
}

//NewConfig
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
	}
}
