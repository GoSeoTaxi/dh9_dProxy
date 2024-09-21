package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	ExternalAPIURLs []string      // Массив URL для внешних API
	GostV3Host      string        // Хост сервера Gost v3
	GostV3Port      string        // Порт сервера Gost v3
	GostV3Username  string        // Логин для Gost v3
	GostV3Password  string        // Пароль для Gost v3
	ProxyServerHost string        // Хост прокси сервера
	ProxyServerPort int           // Порт начального прокси сервера
	MaxProxyServers int           // Максимальное количество прокси-серверов на прокси ноду
	ProxyTTL        time.Duration // Время жизни прокси-сервера
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Не удалось загрузить файл .env, используем системные переменные")
	}

	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		ExternalAPIURLs: parseURLs(getEnv("EXTERNAL_API_URL", "")),
		GostV3Host:      getEnv("GOSTV3_API_HOST", "localhost"),
		GostV3Port:      getEnv("GOSTV3_API_PORT", "8080"),
		GostV3Username:  getEnv("GOSTV3_USERNAME", ""),
		GostV3Password:  getEnv("GOSTV3_PASSWORD", ""),
		ProxyServerHost: getEnv("PROXY_SERVER_HOST", ""),
		ProxyServerPort: getIntEnv("PROXY_SERVER_PORT", 30000),
		MaxProxyServers: getIntEnv("MAX_PROXY_SERVERS", 10),
		ProxyTTL:        getDurationEnv("PROXY_TTL", time.Minute),
	}
}

func parseURLs(urls string) []string {
	if urls == "" {
		return []string{}
	}
	return strings.Split(urls, ";")
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getDurationEnv(key string, defaultVal time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if durationValue, err := time.ParseDuration(value); err == nil {
			return durationValue
		}
	}
	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultVal
}
