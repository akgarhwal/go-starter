package config

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var config appConfig

type appConfig struct {
	AppPort string
	AppKey  string // all off local encryption will use this key
	LogPath string
	// database config
	DbDialeg   string
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	// Redis
	RedisHost     string
	RedisPort     string
	RedisDB       int
	RedisPassword string
	// key
	PrivateKey string
	PublicKey  string
	// jwt
	JwtTokenType      string
	JwtTokenExpired   time.Duration // in second
	JwtRefreshExpired time.Duration // in second
}

func init() {
	var once sync.Once
	once.Do(func() {
		config = load()
	})
}

func load() appConfig {

	jwtTokenExp := "%JWT_TOKEN_EXPIRED%"
	jwtRefreshExp := "%JWT_REFRESH_EXPIRED%"

	jwtTokenDuration, _ := time.ParseDuration(jwtTokenExp)
	jwtRefreshDuration, _ := time.ParseDuration(jwtRefreshExp)

	envRedisDB := "%REDIS_DB%"
	redisDB, err := strconv.Atoi(envRedisDB)
	if err != nil {
		log.Println("Err: Cannot parse redisDB into int")
	}
	return appConfig{
		AppPort: "%APP_PORT%",
		AppKey:  "%APP_KEY%",
		LogPath: "%LOG_PATH%",
		// db configure
		DbDialeg:   "%DB_DIALEG%",
		DbHost:     "%DB_HOST%",
		DbPort:     "%DB_PORT%",
		DbName:     "%DB_NAME%",
		DbUsername: "%DB_USERNAME%",
		DbPassword: "%DB_PASSWORD%",
		// redis
		RedisHost:     "%REDIS_HOST%",
		RedisPort:     "%REDIS_PORT%",
		RedisDB:       redisDB,
		RedisPassword: "%REDIS_PASSWORD%",
		// key
		PrivateKey: "%PRIVATE_KEY%",
		PublicKey:  "%PUBLIC_KEY%",
		// Jwt Configuration
		JwtTokenType:      "Bearer",
		JwtTokenExpired:   jwtTokenDuration,   // in second
		JwtRefreshExpired: jwtRefreshDuration, // in second
	}
}

func Get() appConfig {
	return config
}
