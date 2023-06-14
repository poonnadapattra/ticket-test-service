package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/poonnadapattra/ticket-test-service/internal/config"
	"github.com/poonnadapattra/ticket-test-service/internal/tickets"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Hello")
	config := initConfig()
	db := initDB(config)

	ticketHandler := tickets.Newhandler(tickets.NewService(tickets.NewRepository(db)))

	e := echo.New()
	external := e.Group("/api")

	external.GET("/healt", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	external.GET("/ticket", ticketHandler.GetTicket)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("shutting down the server")
	}
}

func initConfig() (c config.Config) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	viper.AutomaticEnv()
	viper.Unmarshal(&c)

	return
}

func initDB(c config.Config) (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", c.Database.Host, c.Database.Username, c.Database.Password, c.Database.DBName, c.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %v", err))
	}

	return
}
