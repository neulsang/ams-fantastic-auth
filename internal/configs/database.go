package configs

import "ams-fantastic-auth/pkg/env"

type DBConfig struct {
	UserName string
	Password string
	HostName string
	DBName   string
}

func Database() DBConfig {
	userName := env.Get("DATABASE_USER_NAME", "tester")
	password := env.Get("DATABASE_PASSWORD", "test001")
	hostName := env.Get("DATABASE_HOST_NAME", "localhost:3306")
	dbName := env.Get("DATABASE_DB_NAME", "testdb")
	return DBConfig{
		UserName: userName,
		Password: password,
		HostName: hostName,
		DBName:   dbName,
	}
}
