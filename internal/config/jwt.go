package config

import "ams-fantastic-auth/pkg/env"

type JWT struct {
	AccessTokenSecretKey  string //`mapstructure:"JWT_ACCESS_TOKEN_SECRET_KEY"`
	AccessTokenPublicKey  string //`mapstructure:"JWT_ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiredIn  int    //`mapstructure:"JWT_ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxage     int    //`mapstructure:"JWT_ACCESS_TOKEN_MAXAGE"`
	RefreshTokenSecretKey string //`mapstructure:"JWT_REFRESH_TOKEN_SECRET_KEY"`
	RefreshTokenPublicKey string //`mapstructure:"JWT_REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiredIn int    //`mapstructure:"JWT_REFRESH_TOKEN_EXPIRED_IN"`
	RefreshTokenMaxage    int    //`mapstructure:"JWT_REFRESH_TOKEN_MAXAGE"`
}

func (j *JWT) LoadConfig() {
	j.AccessTokenSecretKey = env.ReadAsStr("JWT_ACCESS_TOKEN_SECRET_KEY", "test-access-jwt-secret-key")
	j.AccessTokenPublicKey = env.ReadAsStr("JWT_ACCESS_TOKEN_PUBLIC_KEY", "")
	j.AccessTokenExpiredIn = env.ReadAsInt("JWT_ACCESS_TOKEN_EXPIRED_IN", 15)
	j.AccessTokenMaxage = env.ReadAsInt("JWT_ACCESS_TOKEN_MAXAGE", 15)
	j.RefreshTokenSecretKey = env.ReadAsStr("JWT_REFRESH_TOKEN_SECRET_KEY", "test-refresh-jwt-secret-key")
	j.RefreshTokenPublicKey = env.ReadAsStr("JWT_REFRESH_TOKEN_PUBLIC_KEY", "")
	j.RefreshTokenExpiredIn = env.ReadAsInt("JWT_REFRESH_TOKEN_EXPIRED_IN", 30)
	j.RefreshTokenMaxage = env.ReadAsInt("JWT_REFRESH_TOKEN_MAXAGE", 30)
}
