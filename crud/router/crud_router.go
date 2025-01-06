package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"teach/model"

	"github.com/labstack/echo"
)

func (crud CrudRouters) createData(c echo.Context) error {

	request := new(model.TestData)
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}

	if err := json.Unmarshal(bodyBytes, request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}
	if request.NameData == "" || request.AgeData == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}

	err = crud.CrudService.CreateData(context.TODO(), *request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": *request,
		"msg":  "success",
	})
}

func (crud CrudRouters) readData(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := crud.CrudService.ReadAllData(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	fmt.Println("data = ", data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": data,
	})
}

func (crud CrudRouters) updateData(c echo.Context) error {
	request := new(model.TestData)
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}

	if err := json.Unmarshal(bodyBytes, request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}
	if request.NameData == "" || request.AgeData == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "invalid request body",
		})
	}

	err = crud.CrudService.UpdateData(context.TODO(), request.NameData, request.AgeData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	err = crud.CrudService.UpdateData(context.TODO(), request.NameData, request.AgeData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": *request,
		"msg":  "success",
	})
}

// func (crud CrudRouters) deleteData(c echo.Context) error {
// }
