// main.go
package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string
	Password string
}

func loginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"HighlightColor":  "white",
		"BackgroundColor": "white",
		"TextColor":       "#000",
	})
}

func submitLoginHandler(c echo.Context) error {
	// Retrieve form values
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Perform authentication here (e.g., check against a database)
	// For this example, we'll just do a simple check for demonstration purposes
	validUsername := "user"
	validPassword := "pass"

	if username == validUsername && password == validPassword {
		return c.String(http.StatusOK, "Login successful!")
	}

	return c.String(http.StatusUnauthorized, "Invalid username or password.")
}

func main() {
	e := echo.New()
	e.Static("/assets", "static")
	// Parse templates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Routes
	e.GET("/login", loginHandler)
	e.POST("/login", submitLoginHandler)

	// Start the server
	e.Start(":8081")
}

// Template struct to hold the parsed templates
type Template struct {
	templates *template.Template
}

// Render method to implement the echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
