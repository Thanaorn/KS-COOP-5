package routers

import (
	"net/http"
	"teach/model"

	"github.com/labstack/echo/v4"
)

func (cr ConfigRouters) SetUsers(c echo.Context) error {
	request := new(model.UserInformationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: "Invalid request payload: " + err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}
	if request.UserID == "" {
		return c.JSON(http.StatusBadRequest, model.StatusResponse{
			Message: "UserID is required",
			Status:  http.StatusBadRequest,
		})
	}
	if request.Name == "" {
		return c.JSON(http.StatusBadRequest, model.StatusResponse{
			Message: "Name is required",
			Status:  http.StatusBadRequest,
		})
	}
	if request.Age == "" {
		return c.JSON(http.StatusBadRequest, model.StatusResponse{
			Message: "Age is required",
			Status:  http.StatusBadRequest,
		})
	}
	if request.IDCard == "" {
		return c.JSON(http.StatusBadRequest, model.StatusResponse{
			Message: "IDCard is required",
			Status:  http.StatusBadRequest,
		})
	}

	context := c.Request().Context()
	info := model.InitInformationRedis{
		UserID: request.UserID,
		Name:   request.Name,
		Age:    request.Age,
		Contact: model.ContactRedis{
			Email: request.Contact.Email,
			Phone: request.Contact.Phone,
			Address: model.AddressRedis{
				Street:  request.Contact.Address.Street,
				City:    request.Contact.Address.City,
				Zipcode: request.Contact.Address.Zipcode,
			},
		},
	}

	err := cr.ConfigService.SetUserRedisService(context, info.UserID, info)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.StatusResponse{
			Message: "Failed to set user in Redis: " + err.Error(),
			Status:  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, model.StatusResponse{
		Message: "User information saved successfully",
		Status:  http.StatusOK,
	})
}
func (cr ConfigRouters) GetUsers(c echo.Context) error {
	request := new(model.UserIDInformationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}
	context := c.Request().Context()
	info, err := cr.ConfigService.GetUserRedisService(context, request.UserId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}
	response := model.UserInformationResponse{
		UserID: info.UserID,
		Name:   info.Name,
		Age:    info.Age,
		Contact: model.ContactResponse{
			Email: info.Contact.Email,
			Phone: info.Contact.Phone,
			Address: model.AddressResponse{
				Street:  info.Contact.Address.Street,
				City:    info.Contact.Address.City,
				Zipcode: info.Contact.Address.Zipcode,
			},
		},
	}

	return c.JSON(http.StatusOK, response)

}

func (cr ConfigRouters) DeleteUsers(c echo.Context) error {

	request := new(model.UserIDInformationRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}
	context := c.Request().Context()
	err := cr.ConfigService.DeleteUserRedisService(context, request.UserId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.StatusResponse{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		})
	}

	return c.JSON(http.StatusOK, model.StatusResponse{
		Message: "User information deleted successfully",
		Status:  http.StatusOK,
	})
}
