package models

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	u "golang-api/utils"
)

//Token - model
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//Account - model
type Account struct {
	gorm.Model
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//Validate - model
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if resp, ok := u.CheckValidPhone(account.Phone); !ok {
		return u.Message(false, resp), false
	}

	return u.Message(true, "success"), true

}

//Create - model
func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	if temp, ok := GetAccountByEmail(account.Email); ok {
		if temp != nil {
			return u.Message(false, "Email address already in use by another user.")
		}
	} else {
		return u.Message(false, "Connection error. Please retry")
	}

	if temp, ok := GetAccountByPhone(account.Phone); ok {
		if temp != nil {
			return u.Message(false, "Phone number already in use by another user.")
		}
	} else {
		return u.Message(false, "Connection error. Please retry")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

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
	account, ok := GetAccountByPhone(phone)
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
func GetAccountByEmail(email string) (*Account, bool) {
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
func GetAccountByPhone(phone string) (*Account, bool) {
	acc := &Account{}
	err := GetDB().Table("accounts").Where("phone = ?", phone).First(acc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, true
		}
		return nil, false
	}
	return acc, true
}
