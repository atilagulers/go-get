package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	// logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e.GET("/", func(c echo.Context) error {
		res, err := http.Get("https://api.trendyol.com/grocerygw/categories")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var jsonResponse []map[string]any
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, jsonResponse)
	})

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))

}
