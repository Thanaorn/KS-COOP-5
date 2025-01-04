package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// func (crud CrudRouters) createData(c echo.Context) error {
// }
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

// func (crud CrudRouters) updateData(c echo.Context) error {

// }

// func (crud CrudRouters) deleteData(c echo.Context) error {
// }
