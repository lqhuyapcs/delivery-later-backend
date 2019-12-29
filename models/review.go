package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//Item
type Review struct {
	gorm.Model
	AccountId uint    `json:"account_id"`
	StoreId   uint    `json:"store_id"`
	Rate      float64 `json:"rate"`
	Content   string  `json:"content"`
}

//Create Review
func (review *Review) CreateReview() map[string]interface{} {
	GetDB().AutoMigrate(&Review{})
	GetDB().Create(review)
	if review.ID <= 0 {
		return u.Message(false, "Error when create new review")
	}
	sto := &Store{}
	if temp, ok := getStoreByID(review.StoreId); ok {
		if temp == nil {
			return u.Message(false, "Store doesnt exist !")
		}
		sto = temp
	}
	if sto != nil {
		sto.CalculateRateStore(review.Rate)
		GetDB().Save(sto)
	} else {
		return u.Message(false, "Unknown error")
	}
	response := u.Message(true, "Review has been created")
	response["review"] = review
	return response
}
