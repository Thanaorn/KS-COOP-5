package router

import (
	"net/http"
	ws "teach/internal/webhook/service"

	"github.com/labstack/echo"
)

type WebhookRouters struct {
	WebhookService ws.IWebhookService
}

func NewWebhookRouter(
	e *echo.Echo,
	webhookService ws.IWebhookService,
) {
	wr := WebhookRouters{webhookService}

	Webhook := e.Group("/webhook")

	Webhook.GET("/data", wr.getAllData)

}

func (wr WebhookRouters) getAllData(c echo.Context) error {
	// Your implementation
	err := wr.WebhookService.WebhookService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Data fetched successfully"})
}
