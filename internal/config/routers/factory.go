package routers

import (
	"teach/internal/config/services"
	"teach/internal/httpclient"

	"github.com/labstack/echo"
)

type ConfigRouters struct {
	ConfigService services.IConfigService
	HttpClient    *httpclient.HTTPClient
}

func NewConfigRouters(e *echo.Echo, ss services.IConfigService) {
	cr := ConfigRouters{
		ConfigService: ss,
	}
	gConfig := e.Group("/config")

	gConfig.GET("/users", cr.GetUsers)

}
