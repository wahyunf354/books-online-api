package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RemoveSlashMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
}
