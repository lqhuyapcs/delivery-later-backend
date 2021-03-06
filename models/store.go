package models

import (
	u "golang-api/utils"

	"strings"
	"unicode"

	"github.com/jinzhu/gorm"

	"sort"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//store location
type StoreLocation struct {
	gorm.Model
	StoreId uint    `json:"store_id"`
	Address string  `json:"address"`
	Lat     float64 `gorm:"type:decimal(10,8)"`
	Lng     float64 `gorm:"type:decimal(11,8)"`
}

//Store - model
type Store struct {
	gorm.Model
	Name          string        `json:"name"`
	NameAscii     string        `json:"name_ascii"`
	Owner         string        `json:"owner"`
	AccountId     uint          `json:"account_id"`
	StoreLocation StoreLocation `json:"store_location"`
	Categories    []Category    `gorm:"foreignkey:store_id;association_foreignkey:id"`
	NumberReviews uint          `json:"number_reviews`
	Reviews       []Review      `gorm:"foreignkey:store_id;association_foreignkey:id"`
	City          string        `json:"city"`
	Province      string        `json:"province"`
	Distance      float64       `json:"distance"`
	Rate          float64       `json:"rate"`
}

// main API
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

	GetDB().Create(store)
	GetDB().Create(store.StoreLocation)
	if store.ID <= 0 {
		return u.Message(false, "Error when create new store")
	}

	response := u.Message(true, "Store has been created")
	response["store"] = store
	return response
}

//SearchStoreByName
func SearchStoreByName(name string, Lat float64, Lng float64) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := searchStoreByName(name); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}
		for i := range *temp {
			LatStore := (*temp)[i].StoreLocation.Lat
			LngStore := (*temp)[i].StoreLocation.Lng
			(*temp)[i].Distance = u.Distance(Lat, Lng, LatStore, LngStore)
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
	/*if err, ok := u.CheckValidName(store.Name); !ok {
		// print message if invalid
		return u.Message(false, err)
	}*/

	// check whether store name exists or not
	if temp, ok := getStoreByName(store.Name); ok {
		if temp != nil {
			return u.Message(false, "existed name")
		}
	} else {
		return u.Message(false, "Connection error! Retry later")
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, store.Name)
	store.NameAscii = result
	GetDB().Model(store).Updates(store)
	GetDB().Model(store.StoreLocation).Updates(store.StoreLocation)
	GetDB().Table("stores").Where("ID = ?", store.ID).Preload("StoreLocation").Preload("Categories.Items").First(store)
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

//Get nearest store
func SearchNearestStore(Lat float64, Lng float64) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := getNearestStore(Lat, Lng); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}

		for i := range *temp {
			LatStore := (*temp)[i].StoreLocation.Lat
			LngStore := (*temp)[i].StoreLocation.Lng
			(*temp)[i].Distance = u.Distance(Lat, Lng, LatStore, LngStore)
		}
		sort.SliceStable(*temp, func(i, j int) bool {
			return (*temp)[i].Distance < (*temp)[j].Distance
		})
		response["store"] = temp
	}
	return response
}

//Get highest rate store
func SearchHighestRateStore(Lat float64, Lng float64) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := searchHighestRateStore(); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}
		for i := range *temp {
			LatStore := (*temp)[i].StoreLocation.Lat
			LngStore := (*temp)[i].StoreLocation.Lng
			(*temp)[i].Distance = u.Distance(Lat, Lng, LatStore, LngStore)
		}
		response["store"] = temp
	}
	return response
}

//Get newest store
func SearchNewestStore(Lat float64, Lng float64) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := getNewestStore(); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}
		for i := range *temp {
			LatStore := (*temp)[i].StoreLocation.Lat
			LngStore := (*temp)[i].StoreLocation.Lng
			(*temp)[i].Distance = u.Distance(Lat, Lng, LatStore, LngStore)
		}
		response["store"] = temp
	}
	return response
}

//Support
//Get store by name - model
func getStoreByName(name string) (*Store, bool) {
	sto := &Store{}
	name = strings.ToLower(name)
	err := GetDB().Table("stores").Where("LOWER(name) = ?", name).First(sto).Error
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
	name = strings.ToLower(name)
	err := GetDB().Limit(10).Where("LOWER(name_ascii) LIKE ? OR LOWER(name) LIKE ? ", "%"+name+"%", "%"+name+"%").Preload("StoreLocation").Preload("Categories.Items").Find(sto).Error
	if err != nil {
		if len(*sto) > 0 {
			return sto, true
		}
		return nil, false
	}
	if len(*sto) == 0 {
		return nil, true
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

//Get store by Owner id - model
func getStoreByOwnerID(id uint) (*Store, bool) {
	sto := &Store{}
	err := GetDB().Table("stores").Where("AccountId = ?", id).First(sto).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return sto, true
}

func searchHighestRateStore() (*[]Store, bool) {
	sto := &[]Store{}
	err := GetDB().Table("stores").Order("rate desc").Preload("StoreLocation").Preload("Categories").Preload("Categories.Items").Find(sto)
	if err != nil {
		if len(*sto) > 0 {
			return sto, true
		}
		return nil, false
	}
	if len(*sto) == 0 {
		return nil, true
	}
	return sto, true
}

//Calculate rate
func (store *Store) CalculateRateStore(review_rate float64) {
	if store.NumberReviews == 0 {
		store.Rate = review_rate
		store.NumberReviews += 1
	} else {
		store.NumberReviews += 1
		temp := float64(store.NumberReviews-1) * store.Rate
		store.Rate = (temp + review_rate) / float64(store.NumberReviews)
	}
}

//get nearest store
func getNearestStore(Lat float64, Lng float64) (*[]Store, bool) {
	stolo := &[]StoreLocation{}
	sto := &[]Store{}
	err := GetDB().Table("store_locations").Select("distinct(id), store_id, 2 * 6378.1 * asin(sqrt(sin((radians(lat) - radians($1)) / 2) ^ 2 + cos(radians(lat)) * cos(radians($1)) * sin((radians($2) - radians(lng)) / 2) ^ 2)) as distances ", Lat, Lng).
		Order("distances").Limit(10).Find(stolo).Error
	if err == nil {
		if len(*stolo) > 0 {
			idSlice := []uint{}
			for i := range *stolo {
				idSlice = append(idSlice, (*stolo)[i].StoreId)
			}
			err2 := GetDB().Table("stores").Where("id IN (?)", idSlice).Preload("StoreLocation").Preload("Categories.Items").Find(sto).Error
			if err2 != nil {
				return nil, false
			} else {
				if len(*sto) > 0 {
					return sto, true
				}
				if len(*sto) == 0 {
					return nil, true
				}
			}
		}
		if len(*stolo) == 0 {
			return nil, true
		}
	}
	return nil, false

	/*err := GetDB().
		Select("stores.id, stores.name, stores.number_reviews, stores.rate, store_locations.id, store_locations.store_id, 2 * 6378.1 * asin(sqrt(sin((radians(store_locations.lat) - radians($1)) / 2) ^ 2 + cos(radians(store_locations.lat)) * cos(radians($1)) * sin((radians($2) - radians(store_locations.lng)) / 2) ^ 2)) as distances ", Lat, Lng).
		Joins("JOIN store_locations ON stores.id = store_locations.store_id").
		Group("stores.id, stores.name, stores.number_reviews, stores.rate, store_locations.id, store_locations.store_id, distances").
		Order("").
		Preload("StoreLocation").
		//Preload("Categories").Preload("Categories.Items").Preload("Reviews").
		Limit(10).Find(sto).Error
	if err != nil {
		if len(*sto) > 0 {
			return sto, true
		}
		return nil, false
	}
	if len(*sto) == 0 {
		return nil, true
	}*/
	return sto, true
}

//get newest store
func getNewestStore() (*[]Store, bool) {
	sto := &[]Store{}
	err := GetDB().Table("stores").Order("created_at desc").Preload("StoreLocation").Preload("Categories").Preload("Categories.Items").Find(sto)
	if err != nil {
		if len(*sto) > 0 {
			return sto, true
		}
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

//Get all address
func GetAllAddress() map[string]interface{} {
	response := u.Message(true, "Store location exists")
	if temp, ok := getAllAddress(); ok {
		if temp == nil {
			return u.Message(false, "Store location doesnt exist")
		}
		response["store_location"] = temp
	}
	return response
}
func getAllAddress() (*[]StoreLocation, bool) {
	loca := &[]StoreLocation{}
	err := GetDB().Table("store_locations").Find(loca).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return loca, true
}

//Delete address
func getStoreLocationByID(id uint) (*StoreLocation, bool) {
	loca := &StoreLocation{}
	err := GetDB().Table("store_locations").Where("id = ?", id).First(loca).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return loca, true
}
func (loca *StoreLocation) DeleteAddress() map[string]interface{} {
	if temp, ok := getStoreLocationByID(loca.ID); ok {
		if temp == nil {
			return u.Message(false, "Store location doesnt exist !")
		}
	}
	GetDB().Delete(loca)
	response := u.Message(true, "Store location has been deleted")
	response["store_location"] = nil
	return response
}
