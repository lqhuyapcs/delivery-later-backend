package models

import(
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
)

//Store - model
type Store struct{
	gorm.model
	Name string `json:"storeName valid:"required~Tên không được để trống,runelength(1|50)~Tên không quá 50 kí tự"`
	Location Location
	Owner string `json:"storeOwner" valid:"required~Chủ sở hữu không được để trống,runelength(1|50)~Chủ sở hữu không quá 50 kí tự,alpha~Tên chủ sở hữu chỉ được đặt bằng chữ"`
	Contact string `json:"storeContact" valid:"numeric~Số điện thoại chỉ được đặt bằng số"`
}

//Address - model
type Location struct{
	Address string	`json:"streetName" valid:"required~Địa chỉ không được để trống"`
	City string `json:"city" valid:"-"` 	//Chắc sau cho option chọn tỉnh, thành phố chứ nhìn cái này ngu vl
	Province string	`json:"province" valid:"required~Tỉnh không được để trống,alpha~Tỉnh phải nhập bằng chữ"`
}

//Create - model
func (store *Store) Create() map[string]interface{} {
	if ok, err:= govalidator.ValidateStruct(store); err !=nil {
		return err
	} else {
		// check whether store exists or not
		if temp, ok:=getStoreByName(store.Name); ok{
			if temp!=nil {
				return u.Message(false, "Tên này đã có người sử dụng")
			}
		} else {
			return u.Message(false, "Lỗi kết nối. Vui lòng thử lại sau")
		}
	}
	GetDB().Create(store)

	if store.ID <=0 {
		return u.Message(false, "Đã có lỗi khi tạo tài khoản")
	}
}

//Get store by name - model
func getStoreByName(name string) (*Store, bool) {
	sto :=&Store{}
	err := GetDB().Table("stores").Where("name = ?",name).First(sto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return sto, true
}



















