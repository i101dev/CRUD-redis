package application

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort:   5000,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		fmt.Println()
		fmt.Println("Setting [REDIS_ADDR]")
		fmt.Println()
		cfg.RedisAddress = redisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			fmt.Println()
			fmt.Println("Setting [SERVER_PORT]")
			fmt.Println()
			cfg.ServerPort = uint16(port)
		}
	}

	// fmt.Printf("cfg: %+v\n", cfg)

	return cfg
}
