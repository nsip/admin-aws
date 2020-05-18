package main

// Use api and Test
import (
	"fmt"
	"time"

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

	// Intial update
	doc.UpdateEc2()

	// Regular updates - every 5 minutes
	ticker := time.NewTicker(300 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("TimerHit: update ec2\n")
				doc.UpdateEc2()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	e.GET("/api/ec2s", func(c echo.Context) error {
		return c.JSON(http.StatusOK, doc.Ec2s)
	})

	e.Logger.Fatal(e.Start(":8098"))
}
