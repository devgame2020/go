package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CheckHandler godoc
// @Summary OK 메시지 반환
// @Description check API입니다
// @Tags example
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /chk [get]
func CheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
