package main

import (
	"fmt"
	"net/http"

	"github.com/atilagulers/go-get/internal/scrapers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	products := scrapers.Trendyol()

	for _, product := range products {
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Price: %s\n", product.Price)
		fmt.Printf("Image: %s\n", product.Image)
		fmt.Printf("Url: %s\n", product.Url)
		fmt.Println()
	}

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}

//res, err := http.Get("https://api.trendyol.com/grocerygw/categories")
//		if err != nil {
//			return err
//		}
//		defer res.Body.Close()

//		body, err := io.ReadAll(res.Body)
//		if err != nil {
//			return err
//		}

//		var jsonResponse []map[string]any
//		if err := json.Unmarshal(body, &jsonResponse); err != nil {
//			return err
//		}

//		return c.JSON(http.StatusOK, jsonResponse)
