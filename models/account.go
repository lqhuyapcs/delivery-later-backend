package models

import (
	"os"

	u "golang-api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Token - model
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//Iser location
type AccountLocation struct {
	gorm.Model
	AccountID uint    `json:"account_id"`
	Address   string  `json:"address"`
	Lat       float64 `gorm:"type:decimal(10,8)"`
	Lng       float64 `gorm:"type:decimal(11,8)"`
}

//Account - model
type Account struct {
	gorm.Model
	Name            string          `json:"name"`
	Phone           string          `json:"phone"`
	Email           string          `json:"email"`
	Password        string          `json:"password"`
	Token           string          `json:"token"`
	AccountLocation AccountLocation `json:"account_location"`
	Store           Store
	Orders          []Order  `gorm:"foreignkey:account_id;association_foreignkey:id" json:"orders"`
	Reviews         []Review `gorm:"foreignkey:account_id;association_foreignkey:id"`
}

//Create - model
func (account *Account) Create() map[string]interface{} {

	if err, ok := u.CheckValidMail(account.Email); !ok {
		// print message if invalid
		return u.Message(false, err)
	}
	if err, ok := u.CheckValidPhone(account.Phone); !ok {
		return u.Message(false, err)
	}
	if temp, ok := getAccountByEmail(account.Email); ok {
		if temp != nil {
			return u.Message(false, "Email address already in use by another user.")
		}
	} else {
		return u.Message(false, "Connection error when find email. Please retry")
	}

	if temp, ok := getAccountByPhone(account.Phone); ok {
		if temp != nil {
			return u.Message(false, "Phone number already in use by another user.")
		}
	} else {
		return u.Message(false, "Connection error. Please retry")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)
	GetDB().Create(account.AccountLocation)
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

//Authenticate - model
func Authenticate(phone, password string) map[string]interface{} {

	account := &Account{}
	account, ok := getAccountByPhone(phone)
	if ok {
		if account == nil {
			return u.Message(false, "Phone number not found")
		}
	} else {
		return u.Message(false, "Connection error. Please retry")
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

//GetAccountByEmail - model
func getAccountByEmail(email string) (*Account, bool) {
	acc := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(acc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return acc, true
}

//GetAccountByPhone - model
func getAccountByPhone(phone string) (*Account, bool) {
	acc := &Account{}
	err := GetDB().Table("accounts").Where("phone = ?", phone).Preload("AccountLocation").Preload("Store").Preload("Orders").Preload("Reviews").First(acc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return acc, true
}

//UpdateAccount - model
func (account *Account) UpdateAccount() map[string]interface{} {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	//Create new JWT token for the newly registered account
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password
	GetDB().Model(account).Updates(account)
	GetDB().Model(account.AccountLocation).Updates(account.AccountLocation)
	GetDB().Table("accounts").Where("ID = ?", account.ID).Preload("AccountLocation").Preload("Store").Preload("Orders").Preload("Reviews").First(account)
	response := u.Message(true, "Account has been updated")
	response["account"] = account
	return response
}
