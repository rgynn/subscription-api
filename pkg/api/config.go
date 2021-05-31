package api

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	PTSURL        string
	ClientTimeout time.Duration
	IdleTimeout   time.Duration
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
}

func NewConfigFromEnv() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to get env variables: %w", err)
	}

	host := os.Getenv("HOST")
	if host == "" {
		return nil, errors.New("no HOST env variable set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("no PORT env variable set")
	}

	ptsurl := os.Getenv("PTS_URL")
	if ptsurl == "" {
		return nil, errors.New("no PTS_URL env variable set")
	}

	client, err := time.ParseDuration(os.Getenv("TIMEOUT_CLIENT"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse TIMEOUT_CLIENT env variable to time.Duration: %w", err)
	}

	idle, err := time.ParseDuration(os.Getenv("TIMEOUT_IDLE"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse TIMEOUT_IDLE env variable to time.Duration: %w", err)
	}

	read, err := time.ParseDuration(os.Getenv("TIMEOUT_READ"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse TIMEOUT_READ env variable to time.Duration: %w", err)
	}

	write, err := time.ParseDuration(os.Getenv("TIMEOUT_WRITE"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse TIMEOUT_WRITE env variable to time.Duration: %w", err)
	}

	return &Config{
		Port:          fmt.Sprintf("%s:%s", host, port),
		PTSURL:        ptsurl,
		ClientTimeout: client,
		IdleTimeout:   idle,
		ReadTimeout:   read,
		WriteTimeout:  write,
	}, nil
}
