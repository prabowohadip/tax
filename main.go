package main

import (
	"github.com/prabowohadip/tax/config"
	"github.com/prabowohadip/tax/model"
	"sync"
)


func init() {
	config.GetInstanceDb().AutoMigrate(&model.OrderItemSchema{})
}

func main() {

	defer config.GetInstanceDb().Close()
	wg := sync.WaitGroup{}

	// http handler
	wg.Add(1)
	go func() {
		controller.Routing()
		wg.Done()
	}()

	wg.Wait()
}