package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//OrderItem
type OrderItem struct {
	gorm.Model
	ItemId  uint   `"json:item_id"`
	OrderId uint   `"json:order_id"`
	Amount  uint   `"json:amount"`
	Price   string `"json:price"`
}

//Create OrderItem
func (orderItem *OrderItem) CreateOrderItem() map[string]interface{} {
	GetDB().Create(orderItem)
	if orderItem.ID <= 0 {
		return u.Message(false, "Error when create new orderItem")
	}
	response := u.Message(true, "OrderItem has been created")
	response["orderitem"] = orderItem
	return response
}

//Update OrderItem
func (orderItem *OrderItem) UpdateOrderItem() map[string]interface{} {
	GetDB().Save(orderItem)
	response := u.Message(true, "OrderItem has been updated")
	response["orderitem"] = orderItem
	return response
}

//Delete Orderitem
func (orderItem *OrderItem) DeleteOrderItem() map[string]interface{} {
	GetDB().Delete(orderItem)
	response := u.Message(true, "OrderItem has been deleted")
	response["orderitem"] = nil
	return response
}
