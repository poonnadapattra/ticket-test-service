package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/poonnadapattra/ticket-test-service/internal/config"
	"github.com/poonnadapattra/ticket-test-service/internal/contacts"
	"github.com/poonnadapattra/ticket-test-service/internal/tickets"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config := initConfig()
	db := initDB(config)

	ticketHandler := tickets.Newhandler(tickets.NewService(tickets.NewRepository(db)))
	contactHandler := contacts.Newhandler(contacts.NewService(contacts.NewRepository(db)))

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	external := e.Group("/api")

	external.GET("/healt", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	external.GET("/tickets/status", ticketHandler.GetTicketCount)
	external.GET("/tickets", ticketHandler.GetTicket)
	external.POST("/tickets", ticketHandler.CreateTicket)
	external.PATCH("/tickets", ticketHandler.UpdateTicket)
	external.DELETE("/tickets", ticketHandler.DeleteTicket)

	external.GET("/contacts", contactHandler.GetContact)

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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %v", err))
	}

	return
}
