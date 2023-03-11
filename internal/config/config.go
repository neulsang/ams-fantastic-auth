package config

import (
	"sync"
)

var once sync.Once

type Root struct {
	serverConfig *Server
	loggerConfig *Logger
	rdbConfig    *RDB
	jwtConfig    *JWT
}

func LoadConfig() *Root {
	root := new(Root)
	once.Do(func() {
		serverConfig := new(Server)
		serverConfig.LoadConfig()
		root.serverConfig = serverConfig

		loggerConfig := new(Logger)
		loggerConfig.LoadConfig()
		root.loggerConfig = loggerConfig

		rdbConfig := new(RDB)
		rdbConfig.LoadConfig()
		root.rdbConfig = rdbConfig

		jwtConfig := new(JWT)
		jwtConfig.LoadConfig()
		root.jwtConfig = jwtConfig
	})
	return root
}

func (r *Root) Server() *Server {
	return r.serverConfig
}

func (r *Root) Logger() *Logger {
	return r.loggerConfig
}

func (r *Root) RDB() *RDB {
	return r.rdbConfig
}

func (r *Root) JWT() *JWT {
	return r.jwtConfig
}