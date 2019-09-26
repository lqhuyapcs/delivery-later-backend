package models

import (
	valid "golang-api/validator"
	u "golang-api/utils"
	"gopkg.in/go-playground/validator.v9"
	"github.com/jinzhu/gorm"
)

//Store - model
type Store struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Owner    string `json:"storeOwner"`
	OwnerID string `json:"ownerID";sql:"-"`
	Location []*Location `json:"-"`
}

//Location - model
type Location struct {
	Address  string `json:"streetName" validate:"required"`
	City     string `json:"city" validate:"required"`
	Province string `json:"province" validate:"required"`
}

//Create
func (store *Store) Create() map[string]interface{} {
	//check whether store is valid
	if err,ok := valid.checkValidName(store.Name); !ok {
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
	 }
	 GetDB().Create(store)

	if store.ID <= 0 {
	 	return u.Message(false, "Error when create new store")
	 }  //Code ông ông tự xử nhé :v Tui code mẫu dưới này

	response := u.Message(true, "Store has been created")
	response["store"] = store
	return response
}


//Update
func (store *Store) UpdateStore() map[string]interface{} {
	GetDB().Where("ID = ?", store.ID).First(store)

	// check update valid
	if err,ok := valid.checkValidName(store.Name); !ok {

		 // print message if invalid
		 return u.Message(false, err)
	}

	// check whether store name exists or not
	if temp, ok:=getStoreByName(store.Name); ok{
			if temp!=nil {
				return u.Message(false, "Tên này đã có người sử dụng")
			}
		} else {
			return u.Message(false, "Lỗi kết nối. Vui lòng thử lại sau")
		}
	}
	GetDB().Save(store)
	response := u.Message(true , "Store has been updated")
	response["store"] = store 
	return response
}

//Delete 
func (store *Store) DeleteStore() {
	GetDB().Delete(store)
}

//Get store by name - model
func getStoreByName(name string) (*Store, bool) {
	sto :=&Store{}
	err := GetDB().Table("Store").Where("name = ?",name).First(sto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return sto, true
}



