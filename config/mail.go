package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Sender   string
	Username string
	Password string
}

func LoadSMTP() (*SMTPConfig, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	
	if err != nil {
		return nil, err
	}

	smtpConfig := &SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: port,
		Sender: os.Getenv("SMTP_SENDER"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	return smtpConfig, nil
}
