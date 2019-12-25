package models

import (
	u "golang-api/utils"

	"strings"
	"unicode"

	"github.com/jinzhu/gorm"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	m "math"
)

//store location
type StoreLocation struct {
	gorm.Model
	StoreId uint
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
	City          string        `json:"city"`
	Province      string        `json:"province"`
	Distance      float64       `json:"distance"`
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
	err := GetDB().Limit(10).Where("LOWER(name_ascii) LIKE ? OR LOWER(name) LIKE ? ", "%"+name+"%", "%"+name+"%").Preload("Categories.Items").Find(sto).Error
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

//Get nearest store
func SearchNearestStore(address string, Lat float64, Lng float64) map[string]interface{} {
	response := u.Message(true, "Store exists")
	if temp, ok := getNearestStore(address, Lat, Lng); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist")
		}
		for i := range *temp {
			LatStore := (*temp)[i].StoreLocation.Lat
			LngStore := (*temp)[i].StoreLocation.Lng
			(*temp)[i].Distance = Distance(Lat, Lng, LatStore, LngStore)
		}
		response["store"] = temp
	}
	return response
}

func getNearestStore(address string, Lat float64, Lng float64) (*[]Store, bool) {
	sto := &[]Store{}
	/*err := GetDB().Raw(`SELECT * , 2 * 3961 * asin(sqrt((sin(radians((stoLo.lat - $1) / 2))) ^ 2 + cos(radians(stoLo.lat)) * cos(radians($1)) * (sin(radians(($2 - stoLo.lng) / 2))) ^ 2)) as distance
	FROM Stores as sto, store_locations as stoLo
	WHERE sto.ID = stoLo.Store_id
	ORDER BY distance`, Lat, Lng).Scan(sto)*/
	err := GetDB().Select([]string{"*", "2 * 3961 * asin(sqrt((sin(radians((store_locations.lat - $1) / 2))) ^ 2 + cos(radians(store_locations.lat)) * cos(radians($1)) * (sin(radians(($2 - store_locations.lng) / 2))) ^ 2)) as distances"}, Lat, Lng).
		Where("store_locations.store_id = stores.id").
		Joins("JOIN store_locations ON store_locations.store_id = stores.id").
		Order("distances").Preload("StoreLocation").Preload("Categories").Preload("Categories.Items").Find(sto)
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

//	GROUP BY * HAVING 2 * 3961 * asin(sqrt((sin(radians((stoLo.lat - $1) / 2))) ^ 2 + cos(radians(stoLo.lat)) * cos(radians($1)) * (sin(radians(($2 - stoLo.lng) / 2))) ^ 2))
//Preload("Categories").Preload("Categories.Items").
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * m.Pi / 180
	lo1 = lon1 * m.Pi / 180
	la2 = lat2 * m.Pi / 180
	lo2 = lon2 * m.Pi / 180

	r = 6378.1 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + m.Cos(la1)*m.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * m.Asin(m.Sqrt(h))
}

func hsin(theta float64) float64 {
	return m.Pow(m.Sin(theta/2), 2)
}
