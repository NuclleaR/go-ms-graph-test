package common

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	Port string
	Host string
}

func NewConfig(h, p string) Config {
	return Config{
		Host: h,
		Port: p,
	}
}

func NewConfigFromEnv() Config {
	return Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}

func NewConfigFromFlags() Config {
	host := flag.String("host", "localhost", "host to listen to")
	port := flag.String("port", "", "port to listen to")

	if *port == "" || *host == "" {
		log.Fatal("Port and host must be set")
	}

	flag.Parse()

	return Config{
		Host: *host,
		Port: *port,
	}
}
