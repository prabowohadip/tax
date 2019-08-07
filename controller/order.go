package controller

import (
	"github.com/labstack/echo"
	"github.com/prabowohadip/tax/config"
	"github.com/prabowohadip/tax/logic"
	"github.com/prabowohadip/tax/model"
	"sync"
)

var lgc logic.LogicOrderInterface
var once sync.Once

func init() {
	once.Do(func() {

		lgc = logic.NewOrderLogic()

	})
}

func GetLogic() logic.LogicOrderInterface {
	return lgc
}


func CreateOrder(c echo.Context) error {
	orders := new(model.OrderRequest)
	err, data := config.ParsingValidate(c, orders)
	if err != nil {
		return config.ErrorResponse(c, err)
	}
	e,resp := lgc.CreateData(data)
	if e != nil {
		return config.ErrorResponse(c, e)
	}
	return config.SuccessResponse(c,resp)
}

func ListOrder(c echo.Context) error {
	e,resp := lgc.ReadData()
	if e != nil {
		return config.ErrorResponse(c, e)
	}
	return config.SuccessResponse(c,resp)
}
