package router

import (
	cs "teach/internal/crud/service"

	"github.com/labstack/echo"
)

type CrudRouters struct {
}

func NewCrudRouter(e *echo.Echo, cs cs.ICrudService) {

}
