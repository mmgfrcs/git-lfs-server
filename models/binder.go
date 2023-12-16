package models

import (
	"github.com/labstack/echo/v4"
	"strings"
)

type LFSBinder struct{}

func (b *LFSBinder) Bind(i interface{}, c echo.Context) (err error) {
	req := c.Request()

	if !strings.HasPrefix(req.Header.Get(echo.HeaderContentType), "application/vnd.git-lfs+json") {
		db := new(echo.DefaultBinder)
		if err = db.Bind(i, c); err != nil {
			return
		}
	}

	if err = c.Echo().JSONSerializer.Deserialize(c, i); err != nil {
		return
	}

	return nil
}
