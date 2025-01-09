package router

import (
	cs "teach/crud/service"

	"github.com/labstack/echo"
)

type CrudRouters struct {
	CrudService cs.ICrudService
}

func NewCrudRouter(e *echo.Echo, cs cs.ICrudService) {

	crud := CrudRouters{
		CrudService: cs,
	}

	crudTestApi := e.Group("/test")

	crudTestApi.GET("/readData", crud.readData)
	crudTestApi.GET("/readData/:id", crud.readDataId)
	crudTestApi.POST("/createData", crud.createData)
	crudTestApi.POST("/updateData", crud.updateData)
	crudTestApi.DELETE("/deleteData/:id", crud.deleteData)


}
