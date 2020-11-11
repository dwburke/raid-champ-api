package types

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/dwburke/raid-champ-api/db"
	"github.com/jinzhu/gorm"
)

type ApiUser struct {
	Username string `gorm:"column:username;type:varchar(255);primary_key" json:"username"`
	Password string `gorm:"column:password;type:varchar(128)" json:"password"`
}

func (ApiUser) TableName() string {
	return "api_user"
}

func (u *ApiUser) IsAuthorized(route, method string) (bool, error) {
	provdb := db.Open()

	var access ApiUserAccess

	if err := provdb.Select("*").Where("username = ? and route = ? and method = ?", u.Username, route, method).Find(&access).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (u *ApiUser) SetPassword(password string) error {
	hash := hashAndSalt([]byte(password))
	u.Password = hash

	provdb := db.Open()

	if err := provdb.Save(&u).Error; err != nil {
		return err
	}

	return nil
}

func (u *ApiUser) CheckPassword(password string) bool {
	return comparePasswords(u.Password, []byte(password))
}

// below functions found here... https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
