package repository

import (
	"github.com/prabowohadip/tax/config"
	"github.com/prabowohadip/tax/model"
)

func InsertOrders(params interface{}) (e error, data model.OrderItemSchema) {
	paramater := params.(*model.OrderItemRequest)
	data = model.OrderItemSchema{
		Name:paramater.Name,
		TaxCode:paramater.TaxCode,
		Price:paramater.Price,
	}
	e = config.GetInstanceDb().Save(&data).Error

	return
}

func ReadOrder() (e error, data []model.OrderItemSchema) {
	e = config.GetInstanceDb().Find(&data).Error

	return
}