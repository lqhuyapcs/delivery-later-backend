package models

import (
	u "golang-api/utils"
	"gopkg.in/go-playground/validator.v9"
	"github.com/jinzhu/gorm"
)

//Store - model
type Store struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Location string `json:"location" validate:"required,max=100"`
	Owner    string `json:"storeOwner" validate:"required,min=2,max=50"`
	OwnerID string `json:"ownerID" validate:"-"`
	Location Location 
}

//Location - model
type Location struct {
	Address  string `json:"streetName" validate:"required"`
	City     string `json:"city" validate:"required"` //Chắc sau cho option chọn tỉnh, thành phố chứ nhìn cái này ngu vl
	Province string `json:"province" validate:"required"`
}

//Create - model
func (store *Store) Create() map[string]interface{} {
	valid :=validator.New()
	//check whether store is valid
	 if err := valid.Struct(store); err != nil {
	 	return err
	 } else {
		// if valid, check whether store exists or not
	 	if temp, ok := getStoreByName(store.Name); ok {
	 		if temp != nil {
	 			return u.Message(false, "Tên này đã có người sử dụng")
	 		}
	 	} else {
	 		return u.Message(false, "Lỗi kết nối. Vui lòng thử lại sau")
	 	}
	 }
	 GetDB().Create(store)

	if store.ID <= 0 {
	 	return u.Message(false, "Đã có lỗi khi tạo tài khoản")
	 }  //Code ông ông tự xử nhé :v Tui code mẫu dưới này

	response := u.Message(true, "Cửa hàng đã được tạo mới")
	response["store"] = store
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
