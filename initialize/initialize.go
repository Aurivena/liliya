package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io"
	"liliya/models"
	"os"
	"time"
)

const (
	envFilePath = `.env`
)

var (
	Config = models.Config{}
	Env    = models.Env{}
)

func LoadConfig(configPath string) error {
	file, err := os.ReadFile(configPath)
	if err != nil {
		logrus.Errorf("load config file failed, err:%v", err)
		return err
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		logrus.Errorf("unmarshal config file failed, err:%v", err)
		return err
	}

	return nil
}

func LoadEnv(envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		logrus.Errorf("load env failed, err:%v", err)
		return err
	}
	if err := env.Parse(&Env); err != nil {
		logrus.Errorf("parse env failed, err:%v", err)
		return err
	}
	return nil
}

func LoadConfiguration() error {
	if err := LoadEnv(envFilePath); err != nil {
		return err
	}
	logrus.Info("env variables: loaded")
	if err := LoadConfig(Env.ConfigPath); err != nil {
		return err
	}
	logrus.Info("config file: loaded")

	return nil
}

func RunLogger() error {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{})

	currentTime := time.Now()
	yearMonthDir := fmt.Sprintf("logs/%d-%02d", currentTime.Year(), currentTime.Month())

	err := os.MkdirAll(yearMonthDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка создания папки для логов: %v", err)
	}

	logFile := fmt.Sprintf("%s/%d-%02d-%02d.log", yearMonthDir, currentTime.Year(), currentTime.Month(), currentTime.Day())

	logFileHandle, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла логов: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(os.Stdout, logFileHandle))

	return nil
}
