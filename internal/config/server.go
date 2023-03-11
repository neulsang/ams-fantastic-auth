package config

import "ams-fantastic-auth/pkg/env"

type Server struct {
	Host        string `mapstructure:"SERVER_HOST"`
	Port        int    `mapstructure:"SERVER_PORT"`
	Origin      string `mapstructure:"SERVER_ORIGIN"`
	ReadTimeout int    `mapstructure:"SERVER_READ_TIMEOUT"`
}

func (s *Server) LoadConfig() {
	s.Host = env.ReadAsStr("SERVER_HOST", "localhost")
	s.Port = env.ReadAsInt("SERVER_PORT", 9090)
	s.Origin = env.ReadAsStr("SERVER_ORIGIN", "http://localhost:9090")
	s.ReadTimeout = env.ReadAsInt("SERVER_READ_TIMEOUT", 0)

}
