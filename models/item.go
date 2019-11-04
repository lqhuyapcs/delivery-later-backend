package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Item
type Item struct {
	gorm.Model
	CategoryId uint `json:"category_id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
}

//Create Item
func (item *Item) CreateItem() map[string]interface{} {
	if err, ok :=u.CheckValidName(item.Name); !ok {
		return u.Message(false, err)
	}
	GetDB().Create(item)
	if item.ID <= 0 {
		return u.Message(false, "Error when create new item")
	}
	response := u.Message(true, "Item has been created")
	response["item"] = item
	return response
}

//Update item
func (item *Item) UpdateItem() map[string]interface{} {
	if err,ok:=u.CheckValidName(item.Name); !ok {
		return u.Message(false, err)
	}
	GetDB().Save(item)
	response := u.Message(true, "Item has been updated")
	response["item"] = item
	return response
}

//Delete item
func (item *Item) DeleteItem() map[string]interface{} {
	GetDB().Delete(item)
	response := u.Message(true, "Item has been deleted")
	response["item"] = nil
	return response
}
