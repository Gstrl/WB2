package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port string `yaml:"port"`
}

// MustLoad В реальном проекте лучше читать из переменных окружения
func MustLoad() *Config {
	var cfg Config
	// Получаем текущую директорию
	dir, err := os.Getwd()
	fmt.Println(dir)
	if err != nil {
		fmt.Println("Ошибка получения текущего каталога:", err)

	}

	file, err := os.Open(filepath.Join(dir, "develop/dev11/config/local.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// Читаем yaml
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Ошибка при чнении файла конфигурации")
	}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal("Ошибка при парсинге файла конфигурации")
	}

	return &cfg
}
