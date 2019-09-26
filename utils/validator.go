package main


import (
	"regexp"
	"fmt"
)

var regexMail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
//var regexPhone = regexp.MustCompile(`?:^|[^0-9])(1[34578][0-9]{9})(?:$|[^0-9]`)
// check valid email
func checkValidMail(mail string) (string, bool) {
	if mail==""{
		return "Mail is required", false
	}
	if !regexMail.MatchString(mail) {
		return "Invalid mail type", false
	}
	return "success", true
}

// check valid name
func checkValidName(name string) (string, bool) {
	if name == "" {
		return "Name is required", false
	}
	if len(name)<2 || len(name)>50 {
		return "Length of name is between 2 to 50", false
	} 
	return "success", true
}
// check valid phone  (cant find regex)
func checkValidPhone(phone string) (string, bool) {
	if phone == "" {
		return "Phone is required" ,false
	}
	if !regexPhone.MatchString(phone){
		return "Invalid phone number", false
	}
	return "success", true
}
func main() {
	
}