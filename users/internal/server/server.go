package server

import (
	"errors"
	"fmt"
	"users/internal/config"
)

var (
	ErrNoHostProvided = errors.New("no host provided for server")
	ErrNoPortProvided = errors.New("no port provided for server")
)

func Init() error {
	r := NewRouter()
	err, address := getSrvAddress()

	if err != nil {
		return err
	}

	err = r.Run(address)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func getSrvAddress() (error, string) {
	cfg := config.GetConfig()

	port := cfg.GetString("server.port")

	if port == "" {
		return ErrNoPortProvided, ""
	}

	host := cfg.GetString("server.host")

	if host == "" {
		return ErrNoHostProvided, ""
	}

	return nil, fmt.Sprintf("%s:%s", host, port)
}
