package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

type PyramidResponse struct {
	Word          string `json:"word"`
	IsPyramidWord bool   `json:"isPyramid"`
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	log.SetOutput(e.Logger.Output())

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/pyramid/:word", func(c echo.Context) error {
		word := c.Param("word")
		if word == "" {
			return echo.ErrNotFound
		}

		return c.JSON(http.StatusOK, &PyramidResponse{
			Word:          word,
			IsPyramidWord: IsPyramidWord(word),
		})
	})

	err := e.Start("localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}
