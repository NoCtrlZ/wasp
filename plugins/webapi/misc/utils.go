package misc

import (
	"github.com/labstack/echo"
	"net/http"
)

type SimpleResponse struct {
	Error string `json:"err"`
}

func OkJson(c echo.Context, data interface{}) error {
	return c.JSONPretty(http.StatusOK, data, " ")
}

func OkJsonErr(c echo.Context, err error) error {
	serr := ""
	if err != nil {
		serr = err.Error()
	}
	return OkJson(c, &SimpleResponse{Error: serr})
}
