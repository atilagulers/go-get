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

	scrapers := scrapers.New("laptop")

	scrapers.Scrape()

	trendyolProducts := scrapers.Trendyol.Products

	for _, product := range trendyolProducts {
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Price: %s\n", product.Price)
		fmt.Printf("Image: %s\n", product.Image)
		fmt.Printf("Url: %s\n", product.Url)
		fmt.Println()
	}

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
