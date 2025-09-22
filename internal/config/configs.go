package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/prankevich/MyProject/internal/models"
	"os"
)

var AppSettings models.Config

func ReadSettings() error {

	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("Ощибка загрузки файла .env: %w", err)
	}
	file, err := os.Open("internal/config/config.json")
	if err != nil {
		return fmt.Errorf("Ошибка чтения конфига: %w", err)
	}
	defer file.Close()
	if err = json.NewDecoder(file).Decode(&AppSettings); err != nil {
		return fmt.Errorf("Oшибка парсинга файла конфиг: %w", err)
	}
	return nil
}
