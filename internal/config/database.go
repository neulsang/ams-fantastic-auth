package config

import "ams-fantastic-auth/pkg/env"

type RDB struct {
	Host         string `mapstructure:"RDB_HOST"`
	Port         int    `mapstructure:"RDB_PORT"`
	Username     string `mapstructure:"RDB_USER_NAME"`
	Password     string `mapstructure:"RDB_PASSWORD"`
	DatabaseName string `mapstructure:"RDB_DB_NAME"`
}

func (r *RDB) LoadConfig() {
	r.Host = env.ReadAsStr("RDB_HOST", "localhost")
	r.Port = env.ReadAsInt("RDB_PORT", 3306)
	r.Username = env.ReadAsStr("RDB_USER_NAME", "tester")
	r.Password = env.ReadAsStr("RDB_PASSWORD", "test001")
	r.DatabaseName = env.ReadAsStr("RDB_DB_NAME", "testdb")
}
