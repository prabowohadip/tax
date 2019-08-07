package logic

import (
	"errors"
	"fmt"
	"github.com/prabowohadip/tax/model"
	"github.com/prabowohadip/tax/repository"
)

type TaxCode struct {
	Name 		string
	Refundable  bool
}

var TaxCodeList  = map[string]TaxCode{"1":{Name:"Food & Beverage",Refundable:true},
	"2":{Name:"Tobacco",Refundable:false},
	"3":{Name:"Entertainment",Refundable:false}}

type (
	LogicOrderInterface interface {
		CreateData(params interface{}) (error, interface{})
		ReadData() (error, interface{})
	}
	LogicOrder struct {
	}
)

func NewOrderLogic() LogicOrderInterface {
	return &LogicOrder{}
}

func (l *LogicOrder) CreateData(params interface{}) (error, interface{}) {
	paramater := params.(*model.OrderRequest)
	var data = []model.OrderItemTax{}

	for i,val := range paramater.Order{
		if TaxCodeList[val.TaxCode] == (TaxCode{}) {
			return errors.New(fmt.Sprintf("order_item %d :tax_code not valid",i)), nil
		}else {
			var tax float64
			// calc tax
			if val.TaxCode == "1" {
				tax = (float64(10)/float64(100))*val.Price
			}else if val.TaxCode == "2" {
				tax = float64(2)+((float64(2)/float64(100))*val.Price)
			}else if val.TaxCode == "3" {
				if val.Price < float64(100) {
					tax = float64(0)
				}else {
					tax = (float64(1)/float64(100))*(val.Price-float64(100))
				}
			}

			var itemTax = model.OrderItemTax{
				Name:val.Name,
				TaxCode:val.TaxCode,
				Price:val.Price,
				Type:TaxCodeList[val.TaxCode].Name,
				Refundable:TaxCodeList[val.TaxCode].Refundable,
				Tax:tax,
				Amount:tax+val.Price,
			}
			data = append(data,itemTax)
			if e, _ := repository.InsertOrders(&val);e != nil {
				return e, nil
			}
		}
	}

	return nil, data
}

func (l *LogicOrder) ReadData() (error, interface{}) {
	var list = []model.OrderItemTax{}
	e, data := repository.ReadOrder()
	if e == nil {
		for _,val := range data{
			var tax float64
			// calc tax
			if val.TaxCode == "1" {
				tax = (float64(10)/float64(100))*val.Price
			}else if val.TaxCode == "2" {
				tax = float64(2)+((float64(2)/float64(100))*val.Price)
			}else if val.TaxCode == "3" {
				if val.Price < float64(100) {
					tax = float64(0)
				}else {
					tax = (float64(1)/float64(100))*(val.Price-float64(100))
				}
			}

			var itemTax = model.OrderItemTax{
				Name:val.Name,
				TaxCode:val.TaxCode,
				Price:val.Price,
				Type:TaxCodeList[val.TaxCode].Name,
				Refundable:TaxCodeList[val.TaxCode].Refundable,
				Tax:tax,
				Amount:tax+val.Price,
			}
			list = append(list,itemTax)
		}
	}

	return e, list
}

