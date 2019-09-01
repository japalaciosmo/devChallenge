package main

import (
	"github.com/japalaciosmo/devChallenge/api/app/websocket"
	"github.com/japalaciosmo/devChallenge/api/controllers"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/xid"
	"github.com/subosito/gotenv"
)


var (
	AppName string
	BUILD   string
)

func main() {
	if err := gotenv.Load("env"); err != nil {
		log.Printf("error: %s",err.Error())
	}

	var port=os.Getenv("PORT")
	if port==""{
		port="8080"
	}

	// setup echo
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.RemoveTrailingSlash())
	server.Use(requestID)
	server.Use(middleware.CORS())

	websocket.ThePool = websocket.NewPool()
	go websocket.ThePool.Start()

	// health endpoint
	server.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]string{
			"app":   AppName,
			"build": BUILD,
		})
	})


	// register all other controller endpoints.
	controllers.Routes(server)

	// start the server
	server.Logger.Fatal(server.Start(":" + port))
}


func requestID(next echo.HandlerFunc) echo.HandlerFunc {
	const headerKey = "X-Request-Id"

	return func(c echo.Context) error {
		id := c.Request().Header.Get(headerKey)

		if id == "" {
			id = xid.New().String()
		}

		c.Set(headerKey, id)

		c.Response().Header().Set(headerKey, id)
		return next(c)
	}
}