package env

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

const (
	AppEnv                         = "ENV"
	AppEnvDefault                  = "development"
	ServerPort                     = "SERVER_PORT"
	ServerPortDefault              = "8000"
	ServerReadHeaderTimeout        = "SERVER_READ_HEADER_TIMEOUT"
	ServerReadHeaderTimeoutDefault = 2 * time.Second
	ServerWriteTimeout             = "SERVER_WRITE_TIMEOUT"
	ServerWriteTimeoutDefault      = 15 * time.Second
	ServerIdleTimeout              = "IDLE_TIMEOUT"
	ServerIdleTimeoutDefault       = 60 * time.Second
	ServerShutdownTimeout          = "SERVER_SHUTDOWN_TIMEOUT"
	ServerShutdownTimeoutDefault   = 5 * time.Second
	DbHost                         = "DB_HOST"
	DbHostDefault                  = "localhost"
	DbPort                         = "DB_PORT"
	DbPortDefault                  = "5432"
	DbSSLMode                      = "DB_SSL_MODE"
	DbSSLModeDefault               = "disable"
	DbUser                         = "DB_USER"
	DbUserDefault                  = "postgres"
	DbPassword                     = "DB_PASSWORD"
	DbPasswordDefault              = "silver"
	DbName                         = "DB_NAME"
	DbNameDefault                  = "silver"
	JwtSecret                      = "JWT_SECRET"
	JwtExpiresMin                  = "JWT_EXPIRES_MINUTES"
	JwtExpiresMinDefault           = 60
	TaskEncryptionKey              = "TASK_ENCRYPTION_KEY"
	RabbitMQHost                   = "RABBITMQ_HOST"
	RabbitMQHostDefault            = "localhost"
	RabbitMQPort                   = "RABBITMQ_PORT"
	RabbitMQPortDefault            = "5672"
	RabbitMQUser                   = "RABBITMQ_USER"
	RabbitMQUserDefault            = "silver"
	RabbitMQPassword               = "RABBITMQ_PASSWORD"
	RabbitMQPasswordDefault        = "silver"
)

var App app

type app struct {
	Env                     string
	ServerPort              string
	ServerReadHeaderTimeout time.Duration
	ServerWriteTimeout      time.Duration
	ServerIdleTimeout       time.Duration
	ServerShutdownTimeout   time.Duration
	DbHost                  string
	DbPort                  string
	DbSSLMode               string
	DbUser                  string
	DbPassword              string
	DbName                  string
	JwtSecret               string
	JwtExpiresMinutes       time.Duration
	TaskEncryptionKey       string
	RabbitMQHost            string
	RabbitMQPort            string
	RabbitMQUser            string
	RabbitMQPassword        string
}

func init() {
	_ = godotenv.Load()

	App = app{
		Env:                     GetStr(AppEnv, AppEnvDefault),
		ServerPort:              GetStr(ServerPort, ServerPortDefault),
		ServerReadHeaderTimeout: GetTime(ServerReadHeaderTimeout, ServerReadHeaderTimeoutDefault),
		ServerWriteTimeout:      GetTime(ServerWriteTimeout, ServerWriteTimeoutDefault),
		ServerIdleTimeout:       GetTime(ServerIdleTimeout, ServerIdleTimeoutDefault),
		ServerShutdownTimeout:   GetTime(ServerShutdownTimeout, ServerShutdownTimeoutDefault),
		DbHost:                  GetStr(DbHost, DbHostDefault),
		DbPort:                  GetStr(DbPort, DbPortDefault),
		DbSSLMode:               GetStr(DbSSLMode, DbSSLModeDefault),
		DbUser:                  GetStr(DbUser, DbUserDefault),
		DbPassword:              GetStr(DbPassword, DbPasswordDefault),
		DbName:                  GetStr(DbName, DbNameDefault),
		JwtSecret:               os.Getenv(JwtSecret),
		JwtExpiresMinutes:       GetTime(JwtExpiresMin, JwtExpiresMinDefault),
		TaskEncryptionKey:       os.Getenv(TaskEncryptionKey),
		RabbitMQHost:            GetStr(RabbitMQHost, RabbitMQHostDefault),
		RabbitMQPort:            GetStr(RabbitMQPort, RabbitMQPortDefault),
		RabbitMQUser:            GetStr(RabbitMQUser, RabbitMQUserDefault),
		RabbitMQPassword:        GetStr(RabbitMQPassword, RabbitMQPasswordDefault),
	}
}

// GetStr retrieves a string value from the environment variables.
// key: the name of the environment variable
// defaultValue: the value to return if the environment variable is not set
// returns: the value of the environment variable or the defaultValue
func GetStr(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

// GetTime retrieves a time.Duration value from the environment variables.
// key: the name of the environment variable
// defaultValue: the value to return if the environment variable is not set or cannot be parsed
// returns: the parsed duration or the defaultValue
func GetTime(key string, defaultValue time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(val)
	if err != nil {
		return defaultValue
	}
	return duration
}
