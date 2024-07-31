package main

import (
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/atilagulers/go-get/internal/scrapers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templates, err := getTemplates()
	if err != nil {
		e.Logger.Fatal(err)
	}

	renderer := &Template{
		templates: templates,
	}

	e.Static("/static", "internal/ui/static")
	e.Renderer = renderer

	scrapers := scrapers.New("iphone")

	products := scrapers.Scrape()

	//for _, product := range products {
	//	fmt.Printf("Name: %s\n", product.Name)
	//	fmt.Printf("Price: %s\n", product.Price)
	//	fmt.Printf("Image: %s\n", product.Image)
	//	fmt.Printf("Url: %s\n", product.Url)
	//	fmt.Println()
	//}

	e.GET("/", func(c echo.Context) error {
		data := map[string]any{
			"Title":    "Product List",
			"Products": products,
		}
		return c.Render(http.StatusOK, "base.tmpl", data)
	})

	e.Logger.Info("Server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func getTemplates() (*template.Template, error) {
	partials, _ := filepath.Glob("internal/ui/html/partials/*.tmpl")
	pages, _ := filepath.Glob("internal/ui/html/pages/*.tmpl")
	root, _ := filepath.Glob("internal/ui/html/*.tmpl")
	allFiles := append(partials, pages...)
	allFiles = append(allFiles, root...)
	templates, err := template.ParseFiles(allFiles...)
	if err != nil {
		return nil, err
	}
	return templates, nil
}
