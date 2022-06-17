package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	return &Config{
		DBUsername: DBUsername,
		DBPassword: DBPassword,
		DBHost:     DBHost,
		DBName:     DBName,
		DBPort:     DBPort,
	}
}

func MySql(c *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully to connect in database")
	return db, nil
}
