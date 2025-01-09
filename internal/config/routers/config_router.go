package routers

import (
	"net/http"
	"teach/model"

	"github.com/labstack/echo"
)

func (cr ConfigRouters) GetUsers(c echo.Context) error {
	request := new(model.UserInformationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}
	//context := c.Request().Context()

	return nil

}
