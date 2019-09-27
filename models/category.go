package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Category
type Category struct {
	gorm.Model
	Name string `json:"name"`
	Item []Item `gorm:"foreignkey:Categoryrefer`
}

//Item
type Item struct {
	gorm.Model
	Name  string `json:"name"`
	Price string `json:"price"`
}

//Create
func (category *Category) Create() map[string]interface{} {
	if err, ok := u.CheckValidName(category.Name); !ok {
		return u.Message(false, err)
	}
	GetDB().Create(category)

	if category.ID <= 0 {
		return u.Message(false, "Error when create new category")
	}
	response := u.Message(true, "Category has been created")
	response["category"] = category
	return response
}
