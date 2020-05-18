package main

// Use api and Test
import (
	"fmt"

	// "encoding/xml"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	adminaws "github.com/nsip/admin-aws"
)

func main() {
	fmt.Println("Starting ADMIN::AWS")

	doc := adminaws.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	})) // allow cors requests during testing

	e.Static("/", "static")

	e.GET("/api/ec2", func(c echo.Context) error {
		doc.UpdateEc2()
		return c.JSON(http.StatusOK, doc.Ec2s)
	})

	e.Logger.Fatal(e.Start(":8098"))
}
