package varconfig

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnv(key, defaultValue string) string {
	value := LoadEnv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// load only once
func init() {
	dir, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(dir, "env"))
	if err != nil {
		log.Fatalf("Cannot load env. %v", err)
	}
}

func LoadEnv(key string) string {
	return os.Getenv(key)
}

var (
	LogLevel          = GetEnv("LOG_LEVEL", "DEBUG")
	PostgresHost      = GetEnv("POSTGRES_HOST", "127.0.0.1")
	PostgresPort      = GetEnv("POSTGRES_PORT", "5432")
	PostgresUser      = GetEnv("POSTGRES_USER", "postgres")
	PostgresPasswd    = GetEnv("POSTGRES_PASSWD", "admin123")
	PostgresDb        = GetEnv("POSTGRES_DB", "logs")
	RabbitmqURL       = GetEnv("RABBITMQ_URL", "")
	RabbitmqQueue     = "logs"
	ServerPort        = GetEnv("SERVER_PORT", "5010")
	GrpcServerAddress = GetEnv("GRPC_SERVER_ADDRESS", "0.0.0.0:5011")
)
