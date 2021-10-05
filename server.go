package main

import (
	"counter-api/config"
	"counter-api/domain"
	"counter-api/handler"
	"counter-api/repository"
	"counter-api/service"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Server struct {
	e *echo.Echo
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	e := echo.New()
	e.HideBanner = true
	server := &Server{e: e, config: config}

	return server
}

func initializedData() domain.Number {
	jsonFile, err := os.Open("data/data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var number domain.Number
	json.Unmarshal(byteValue, &number)

	return number
}

func (s *Server)Start() error {
	s.e.Server.Addr = fmt.Sprintf(":%d", s.config.Server.Port)
	number := initializedData()
	fmt.Println("number : " , number.Number)
	counterRepository := repository.NewCounterRepository(&number)
	counterService := service.NewCounterService(counterRepository)
	handler.NewCounterHandler(s.e, counterService)

	s.e.GET("/health", s.healthCheck)

	return graceful.ListenAndServe(s.e.Server, time.Second * 10)
}

func (s *Server)healthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}