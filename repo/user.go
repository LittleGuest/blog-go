package repo

import (
	"blog/model"
	"blog/repo/db"
)

// GetUserByAccount get user info by account
func GetUserByAccount(account string) *model.User {
	user := &model.User{}
	db.GetMDB().Where("account = ?", account).First(user)
	return user
}

// VerifyUser verify user
func VerifyUser(account string, email string) bool {
	user := GetCurrentUser()
	if user == nil {
		return false
	}
	if user.Account == account && user.Email == email {
		return true
	}
	return false
}

// GetCurrentUser get current user
func GetCurrentUser() *model.User {
	user := &model.User{}
	db.GetMDB().First(user)
	return user
}

// UpdateUser update user
func UpdateUser(user *model.User) {
	db.GetMDB().Save(user)
}
