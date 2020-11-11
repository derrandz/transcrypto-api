package daemon

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

type (
	DaemonConfig struct {
		Port       string
		PrivateKey string
	}
)

func GetDaemonConfig() *DaemonConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, port_present := os.LookupEnv("PORT")
	if !port_present {
		port = "8080"
		fmt.Println("env value PORT is not present, going with default", port)
	}

	private_key, pk_present := os.LookupEnv("DAEMON_PRIVATE_KEY")
	if !pk_present {
		private_key = "123123123"
		fmt.Println("env value DAEMON_PRIVATE_KEY is not present, going with default", private_key)
	}

	fmt.Println("env", port, private_key)
	return &DaemonConfig{
		Port:       port,
		PrivateKey: private_key,
	}
}
