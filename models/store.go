package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Store - model
type Store struct {
	gorm.Model
	// Name     string `json:"storeName valid:"required~Tên không được để trống,runelength(1|50)~Tên không quá 50 kí tự"`
	// Location Location
	// Owner    string `json:"storeOwner" valid:"required~Chủ sở hữu không được để trống,runelength(1|50)~Chủ sở hữu không quá 50 kí tự,alpha~Tên chủ sở hữu chỉ được đặt bằng chữ"`
	// Contact  string `json:"storeContact" valid:"numeric~Số điện thoại chỉ được đặt bằng số"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Onwer    string `json:"storeOwner"`
	Contact  string `json:"storeContact"`
}

//Location - model
type Location struct {
	Address  string `json:"streetName" valid:"required~Địa chỉ không được để trống"`
	City     string `json:"city" valid:"-"` //Chắc sau cho option chọn tỉnh, thành phố chứ nhìn cái này ngu vl
	Province string `json:"province" valid:"required~Tỉnh không được để trống,alpha~Tỉnh phải nhập bằng chữ"`
}

//Create - model
func (store *Store) Create() map[string]interface{} {
	// if ok, err := govalidator.ValidateStruct(store); err != nil {
	// 	return err
	// } else {
	// 	// check whether store exists or not
	// 	if temp, ok := getStoreByName(store.Name); ok {
	// 		if temp != nil {
	// 			return u.Message(false, "Tên này đã có người sử dụng")
	// 		}
	// 	} else {
	// 		return u.Message(false, "Lỗi kết nối. Vui lòng thử lại sau")
	// 	}
	// }
	// GetDB().Create(store)

	// if store.ID <= 0 {
	// 	return u.Message(false, "Đã có lỗi khi tạo tài khoản")
	// }  //Code ông ông tự xử nhé :v Tui code mẫu dưới này

	if temp, ok := getStoreByName(store.Name); ok {
		if temp != nil {
			return u.Message(false, "Tên cửa hàng đã tồn tại")
		}
	} else {
		return u.Message(false, "Connection error. Please retry")
	}
	GetDB().Create(store)
	if store.ID <= 0 {
		return u.Message(false, "Không tạo được cửa hàng, lỗi kết nối!")
	}
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
