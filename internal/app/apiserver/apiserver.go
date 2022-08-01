package apiserver

//APISERVER
type APIserver struct {
	config *Config
}

//NEW
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
	}
}

//START
func (s *APIserver) Start() error {
	return nil
}
