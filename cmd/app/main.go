package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))

}
