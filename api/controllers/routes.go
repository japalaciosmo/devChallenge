package controllers

import (
	"fmt"
	"github.com/japalaciosmo/devChallenge/api/app/common"
	"github.com/japalaciosmo/devChallenge/api/controllers/websocket"
	"github.com/labstack/echo"


)

// noop no operation function used for coding stubbing.
func noop(ctx echo.Context) error {
	return app.Error(ctx, fmt.Errorf("endpoint not yet implemented"))
}

// Routes mounts all controller endpoints on the server.
func Routes(e *echo.Echo) {
	// core routes
	WebsocketRoutes(e)

}


// ConnectionRoutes endpoints for managing social connection records.
func WebsocketRoutes(e *echo.Echo) {
	g := e.Group("/ws")
	g.GET("", websocket.StartWs)
}