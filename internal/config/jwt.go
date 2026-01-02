package config

import "time"

func JWTSecret() string {
	return Get("JWT_SECRET")
}

func JWTExpire() time.Duration {
	return time.Hour * 24
}
