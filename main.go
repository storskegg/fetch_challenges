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

type VersionsResponse struct {
	VersionA     string `json:"versionA"`
	VersionB     string `json:"versionB"`
	Relationship string `json:"relationship"`
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

		return c.JSONPretty(http.StatusOK, &PyramidResponse{
			Word:          word,
			IsPyramidWord: IsPyramidWord(word),
		}, "  ")
	})
	apiV1.GET("/versions/:versionA/:versionB", func(c echo.Context) error {
		versionA := c.Param("versionA")
		versionB := c.Param("versionB")

		if versionA == "" || versionB == "" {
			return echo.ErrNotFound
		}

		relationship, err := VersionStringCompare(versionA, versionB)
		if err != nil {
			bad := echo.ErrBadRequest
			bad.Message = "one or both version strings was malformed. ensure they contain only digits and periods"
			return bad
		}

		return c.JSONPretty(http.StatusOK, &VersionsResponse{
			VersionA:     versionA,
			VersionB:     versionB,
			Relationship: relationship,
		}, "  ")
	})

	err := e.Start("localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
}
