package helper

import (
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	Host        string
	Port        string
	JwtSecret   string
	ServiceName string
	Db          struct {
		Host     string
		User     string
		Password string
		Name     string
		Timezone string
	}
	Cors struct {
		AllowedOrigins []string
		AllowedMethods []string
		AllowedHeaders []string
	}
	Swagger struct {
		Protocol string
		Host     string
		Port     string
		BasePath string
	}
	Retry struct {
		Count    int
		Delay    int
		DelayMax int
	}
}

var env = new(Env)

func InitEnv(filenames ...string) error {
	err := godotenv.Load(filenames...)

	if err != nil {
		return err
	}

	env.Host = FromEnv("HOST", "0.0.0.0")
	env.Port = FromEnv("PORT", "7001")
	env.JwtSecret = FromEnv("JWT_SECRET", "")
	env.ServiceName = FromEnv("SERVICE_NAME", "Company Service")
	env.Db.Host = FromEnv("DB_HOST", "")
	env.Db.User = FromEnv("DB_USER", "")
	env.Db.Password = FromEnv("DB_PASSWORD", "")
	env.Db.Name = FromEnv("DB_NAME", "")
	env.Db.Timezone = FromEnv("DB_TIMEZONE", "Local")

	if os.Getenv("DB_TIMEZONE") != "" {
		env.Db.Timezone = url.QueryEscape(os.Getenv("DB_TIMEZONE"))
	}

	env.Cors.AllowedOrigins = []string{"http://localhost:" + env.Port, "http://0.0.0.0:" + env.Port}

	if os.Getenv("CORS_ORIGINS") != "" {
		env.Cors.AllowedOrigins = append(env.Cors.AllowedOrigins, strings.Split(os.Getenv("CORS_ORIGINS"), ",")...)
	}

	env.Cors.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	if os.Getenv("CORS_METHODS") != "" {
		env.Cors.AllowedOrigins = strings.Split(os.Getenv("CORS_METHODS"), ",")
	}

	env.Cors.AllowedHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}

	if os.Getenv("CORS_HEADERS") != "" {
		env.Cors.AllowedHeaders = strings.Split(os.Getenv("CORS_HEADERS"), ",")
	}

	// Swagger
	env.Swagger.Protocol = FromEnv("SWAGGER_PROTOCOL", "http://")
	env.Swagger.Host = FromEnv("SWAGGER_HOST", "localhost")
	env.Swagger.Port = FromEnv("SWAGGER_PORT", env.Port)
	env.Swagger.BasePath = FromEnv("SWAGGER_BASE_PATH", "")

	// Retry
	env.Retry.Count = 0
	if os.Getenv("RETRY_COUNT") != "" {
		env.Retry.Count, _ = strconv.Atoi(os.Getenv("RETRY_COUNT"))
	}

	env.Retry.Delay = 0
	if os.Getenv("RETRY_DELAY") != "" {
		env.Retry.Delay, _ = strconv.Atoi(os.Getenv("RETRY_DELAY"))
	}

	env.Retry.DelayMax = 0
	if os.Getenv("RETRY_DELAY_MAX") != "" {
		env.Retry.DelayMax, _ = strconv.Atoi(os.Getenv("RETRY_DELAY_MAX"))
	}

	return nil
}

func GetEnv() *Env {
	return env
}

func FromEnv(key, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}
