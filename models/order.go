package models

import (
	u "golang-api/utils"

	"time"

	"github.com/jinzhu/gorm"
)

//Order
type Order struct {
	gorm.Model
	AccountId     uint
	StoreId       uint
	Created       string `json:"created"`
	Deadline      string `json:"deadline"`
	OrderDate     time.Time
	OrderDeadline time.Time
	Address       string      `json:"address"`
	Cancel        bool        `json:"cancel"`
	Delivered     bool        `json:"delivered"`
	OrderItems    []OrderItem `gorm:"foreignkey:order_id;association_foreignkey:id" json:"orderitems"`
}

//Create Order
func CreateOrder(order Order) map[string]interface{} {
	temp, err := time.Parse("2006-01-02T15:04:05-07:00", order.Created)
	if err == nil {
		order.OrderDate = temp
	}
	temp, err = time.Parse("2006-01-02T15:04:05-07:00", order.Deadline)
	if err == nil {
		order.OrderDeadline = temp
	}
	GetDB().Create(&order)
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

//Search Completed Order
func SearchCompletedOrder(ID uint) map[string]interface{} {
	response := u.Message(true, "Orders exists")
	if temp, ok := getCompletedOrder(ID); ok {
		if temp == nil {
			return u.Message(false, "Orders doesnt exist")
		}
		response["order"] = temp
	}
	return response
}

//Search Incompleted Order
func SearchIncompletedOrder(ID uint) map[string]interface{} {
	response := u.Message(true, "Orders exists")
	if temp, ok := getIncompletedOrder(ID); ok {
		if temp == nil {
			return u.Message(false, "Orders doesnt exist")
		}
		response["order"] = temp
	}
	return response
}

//support
//get completed order
func getCompletedOrder(ID uint) (*[]Order, bool) {
	order := &[]Order{}
	err := GetDB().Table("orders").Where("account_id = ? AND Delivered = ?", ID, true).Order("order_deadline desc").Preload("OrderItems").Find(order).Error
	if err != nil {
		if len(*order) > 0 {
			return order, true
		}
		return nil, false
	}
	if len(*order) == 0 {
		return nil, true
	}
	return order, true
}

//get incompleted order
func getIncompletedOrder(ID uint) (*[]Order, bool) {
	order := &[]Order{}
	err := GetDB().Table("orders").Where("account_id = ? AND Delivered = ?", ID, false).Order("order_date desc").Preload("OrderItems").Find(order).Error
	if err != nil {
		if len(*order) > 0 {
			return order, true
		}
		return nil, false
	}
	if len(*order) == 0 {
		return nil, true
	}
	return order, true
}

//get list date of account
func GetListDateOfAccount(ID uint) (*[]Order, bool) {
	order := &[]Order{}
	err := GetDB().Table("orders").Where("account_id = ? AND Delivered = ? AND  DATE(order_deadline) = ? as date", ID, false, "2020/01/02").Order("order_date desc").Preload("OrderItems").Group("date").Find(order).Error
	if err != nil {
		if len(*order) > 0 {
			return order, true
		}
		return nil, false
	}
	if len(*order) == 0 {
		return nil, true
	}
	return order, true
}
