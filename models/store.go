package models

import (
	u "golang-api/utils"

	"unicode"

	"github.com/jinzhu/gorm"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//Store - model
type Store struct {
	gorm.Model
	Name       string     `json:"name"`
	NameAscii  string     `json:"name_ascii"`
	Owner      string     `json:"owner"`
	AccountId  uint       `json:"account_id`
	Categories []Category `gorm:"foreignkey:store_id;association_foreignkey:id" json:"categories"`
	Address    string     `json:"address"`
	City       string     `json:"city"`
	Province   string     `json:"province"`
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
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, store.Name)
	store.NameAscii = result
	GetDB().AutoMigrate(&Store{}, &Category{}, &Item{})
	GetDB().Create(store)

	if store.ID <= 0 {
		return u.Message(false, "Error when create new store")
	}

	response := u.Message(true, "Store has been created")
	response["store"] = store
	return response
}

//SearchStoreByName
func SearchStoreByName(name string) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := searchStoreByName(name); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}
		response["store"] = temp
	}
	return response
}

//UpdateStore - Update
func (store *Store) UpdateStore() map[string]interface{} {

	if temp, ok := getStoreByID(store.ID); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist !")
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
			return u.Message(false, "existed name")
		}
	} else {
		return u.Message(false, "Connection error! Retry later")
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
			return u.Message(false, "Store doesnt exist !")
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

//Query store by name - model
func searchStoreByName(name string) (*[]Store, bool) {
	sto := &[]Store{}
	err := GetDB().Table("stores").Where("name_ascii LIKE ?", "%"+name+"%").Find(sto).Error
	if err != nil {
		return nil, false
	}
	if len(*sto) == 0 {
		return nil, true
	}
	return sto, true
}

// normalize name
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
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
