package websocket

import (
	"fmt"
	"github.com/japalaciosmo/devChallenge/api/app/common"

	"github.com/labstack/echo"

	"github.com/japalaciosmo/devChallenge/api/app/websocket"
)

func StartWs(ctx echo.Context) error {


	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(ctx.Response(), ctx.Request())
	if err != nil {
		return app.BadRequest(ctx, err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: websocket.ThePool,
	}

	websocket.ThePool.Register <- client
	client.Read()

	return app.Success(ctx, nil)
}


