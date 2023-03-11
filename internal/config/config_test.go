package config_test

import (
	"ams-fantastic-auth/internal/config"
	"log"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := config.LoadConfig()
	log.Println("server: ", cfg.Server())
	log.Println("logger: ", cfg.Logger())
	log.Println("rdb: ", cfg.RDB())
	log.Println("jwt: ", cfg.JWT())
}
