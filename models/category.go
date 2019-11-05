package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Category
type Category struct {
	gorm.Model
	StoreId uint   `json:"store_id"`
	Name    string `json:"name"`
	Items   []Item `gorm:"foreignkey:category_id;association_foreignkey:id" json:"items"`
}

//Create category
func (category *Category) CreateCategory() map[string]interface{} {

	if err, ok := u.CheckValidName(category.Name); !ok {
		return u.Message(false, err)
	}
	GetDB().Create(category)
	//GetDB().Preload("Items").Find(category)
	GetDB().Set("gorm:auto_preload", true).Find(category)
	if category.ID <= 0 {
		return u.Message(false, "Error when create new category")
	}
	db.AutoMigrate(&Category{}, &Item{})
	GetDB().Save(category)
	response := u.Message(true, "Category has been created")
	response["category"] = category
	return response
}

//Update category
func (category *Category) UpdateCategory() map[string]interface{} {
	if err, ok := u.CheckValidName(category.Name); !ok {
		return u.Message(false, err)
	}
	GetDB().Save(category)
	response := u.Message(true, "Category has been updated")
	response["category"] = category
	return response
}

//Delete category
func (category *Category) DeleteCategory() map[string]interface{} {
	GetDB().Delete(category)
	response := u.Message(true, "Store has been deleted")
	response["category"] = nil
	return response
}

// GORM will auto save associations and its reference when creating/updating a record

//Get all category and items
func getAllCategoryItems() (*Category, bool) {
	category := &Category{}
	err := GetDB().Table("categories").First(category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return category, true
}

func getAllCategoryByStore(store *Store) (*Category, bool) {
	category := &Category{}
	err := GetDB().Table("categories").Where("StoreId = ?", store.ID).First(category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return category, true
}
