package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Category
type (
	Category struct {
		gorm.Model
		StoreId uint   `json:"storeid"`
		Name    string `json:"name"`
		Item    []Item `json:"items"`
	}
	//Item

	Item struct {
		gorm.Model
		Name  string `json:"name"`
		Price string `json:"price"`
	}
)

//Create category
func (category *Category) CreateCategory(store *Store) map[string]interface{} {
	if err, ok := u.CheckValidName(category.Name); !ok {
		return u.Message(false, err)
	}
	// category belong to store
	category.StoreId = store.ID
	GetDB().Create(category)
	if category.ID <= 0 {
		return u.Message(false, "Error when create new category")
	}
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
/*Create item
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
}*/

//Delete item
func (item *Item) DeleteStore() map[string]interface{} {
	GetDB().Delete(item)
	response := u.Message(true, "Item has been deleted")
	response["item"] = nil
	return response
}

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
