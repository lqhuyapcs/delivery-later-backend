package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Store - model
type Store struct {
	gorm.Model
	Name     string    `json:"name"`
	Owner    string    `json:"storeOwner"`
	Location *Location `json:"location"`
}

//Location - model
type Location struct {
	Address  string `json:"streetName" validate:"required"`
	City     string `json:"city" validate:"required"`
	Province string `json:"province" validate:"required"`
}

//Create - New Store
func (store *Store) Create() map[string]interface{} {
	//check whether store is valid
	if err, ok := u.CheckValidName(store.Name); !ok {
		// print message if invalid
		return u.Message(false, err)
	}
	// if valid, check whether store exists or not
	if temp, ok := getStoreByName(store.Name); ok {
		if temp != nil {
			return u.Message(false, "Exist name")
		}
	} else {
		return u.Message(false, "Connection error")
	}
	GetDB().Create(store)

	if store.ID <= 0 {
		return u.Message(false, "Error when create new store")
	}

	response := u.Message(true, "Store has been created")
	response["store"] = store
	return response
}

//UpdateStore - Update
func (store *Store) UpdateStore() map[string]interface{} {

	if temp, ok := getStoreByID(store.ID); ok {
		if temp == nil {
			return u.Message(false, "Cửa hàng không tồn tại !")
		}
	}

	// check update valid
	if err, ok := u.CheckValidName(store.Name); !ok {

		// print message if invalid
		return u.Message(false, err)
	}

	// check whether store name exists or not
	if temp, ok := getStoreByName(store.Name); ok {
		if temp != nil {
			return u.Message(false, "Tên này đã có người sử dụng")
		}
	} else {
		return u.Message(false, "Lỗi kết nối. Vui lòng thử lại sau")
	}

	GetDB().Save(store)
	response := u.Message(true, "Store has been updated")
	response["store"] = store
	return response
}

//DeleteStore - Del store
func (store *Store) DeleteStore() map[string]interface{} {
	if temp, ok := getStoreByID(store.ID); ok {
		if temp == nil {
			return u.Message(false, "Cửa hàng không tồn tại !")
		}
	}
	GetDB().Delete(store)
	response := u.Message(true, "Store has been deleted")
	response["store"] = nil
	return response
}

//Get store by name - model
func getStoreByName(name string) (*Store, bool) {
	sto := &Store{}
	err := GetDB().Table("stores").Where("name = ?", name).First(sto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return sto, true
}

//Get store by id - model
func getStoreByID(id uint) (*Store, bool) {
	sto := &Store{}
	err := GetDB().Table("stores").Where("id = ?", id).First(sto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return sto, true
}
