package models

import (
	u "golang-api/utils"

	"time"

	"github.com/jinzhu/gorm"
)

//Order
type Order struct {
	gorm.Model
	AccountId     uint        `"json:account_id"`
	OrderDate     time.Time   `"json:created"`
	OrderDeadline time.Time   `"json:deadline"`
	Status        string      `"json:status"`
	OrderItems    []OrderItem `gorm:"foreignkey:order_id;association_foreignkey:id" json:"orderitems"`
}

//Create Order
func (order *Order) CreateOrder() map[string]interface{} {
	GetDB().AutoMigrate(&Order{}, &OrderItem{})
	GetDB().Create(order)
	if order.ID <= 0 {
		return u.Message(false, "Error when create new order")
	}
	response := u.Message(true, "Order has been created")
	response["order"] = order
	return response
}

//Update Order
func (order *Order) UpdateOrder() map[string]interface{} {
	GetDB().Save(order)
	response := u.Message(true, "Order has been updated")
	response["order"] = order
	return response
}

//Delete Order
func (order *Order) DeleteOrder() map[string]interface{} {
	GetDB().Delete(order)
	response := u.Message(true, "Order has been deleted")
	response["order"] = nil
	return response
}
