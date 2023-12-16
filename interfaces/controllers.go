package interfaces

import "github.com/labstack/echo/v4"

type APIController interface {
	Ping(ctx echo.Context) error
}

type LFSController interface {
	Batch(ctx echo.Context) error
}
