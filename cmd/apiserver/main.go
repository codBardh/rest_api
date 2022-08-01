package apiserver

import (
	"log"
	"rest_api_russian/internal/app/apiserver"
)

var (
	configPath string
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
