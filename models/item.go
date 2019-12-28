package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Item
type Item struct {
	gorm.Model
	CategoryId  uint        `json:"category_id"`
	Name        string      `json:"name"`
	Price       float64     `json:"price"`
	Description string      `json:"description"`
	OrderItems  []OrderItem `gorm:"foreignkey:item_id;association_foreignkey:id" json:"orderitems"`
}

//Create Item
func (item *Item) CreateItem() map[string]interface{} {
	if err, ok := u.CheckValidName(item.Name); !ok {
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
	if err, ok := u.CheckValidName(item.Name); !ok {
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

//Get item by ID
func GetItemByID(id uint) (*Item, bool) {
	item := &Item{}
	err := GetDB().Table("Items").Where("ID = ?", id).First(item).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return item, true
}
