package routers

import (
	"teach/internal/config/services"
	"teach/internal/httpclient"
	"teach/storage"

	"github.com/labstack/echo/v4"
)

type ConfigRouters struct {
	ConfigService services.IConfigService
	HttpClient    *httpclient.HTTPClient
	RedisStorage  *storage.RedisStorage
}

func NewConfigRouters(e *echo.Echo, ss services.IConfigService) {
	cr := ConfigRouters{
		ConfigService: ss,
	}
	gConfig := e.Group("/config")

	gConfig.GET("/users", cr.GetUsers)
	gConfig.POST("/set/users", cr.SetUsers)

}
