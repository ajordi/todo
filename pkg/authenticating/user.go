package authenticating

import "github.com/jinzhu/gorm"

// User defines the properties of a user to be authenticated
type User struct {
	gorm.Model
	UserName  string
	FirstName string
	LastName  string
	Password string
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
