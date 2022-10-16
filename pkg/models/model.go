package model

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Nin      int       `json:"Nin"`
	Name     string    `json:"Name"`
	Address  string    `json:"Address"`
	Password string    `json:"Password"`
	Accounts []Account `json:"Accounts,omitempty" gorm:"foreignKey:ID;references:Bank"`
}

// Account Model
type Account struct {
	gorm.Model
	Bank           string  `json:"Bank"`
	AccountType    string  `json:"AccountType"`
	AccountNumber  string  `json:"AccountNumber"`
	AccountBalance float32 `json:"AccountBalance"`
}

// Creates a new user in the DB.
func (newuser *User) CreateUser(db *gorm.DB) *User {
	db.Create(newuser)
	return newuser
}

// creates a new account in the DB.
func (account *Account) CreateAccount(db *gorm.DB) *Account {
	db.Create(account)
	return account
}

// Return all users in the DB.
func FindUsers(db *gorm.DB) []User {
	var AllUsers []User
	db.Find(&AllUsers)
	return AllUsers
}

// Return all Accounts in the DB.
func FindAccounts(db *gorm.DB) []Account {
	var AllAccounts []Account
	db.Find(&AllAccounts)
	return AllAccounts
}

// Return one users with the NIN provided from the DB.
func FindUser(db *gorm.DB, nin int) (user User) {
	var FoundUser User
	db.Where("Nin=?", nin).Find(&FoundUser)
	return FoundUser
}

// Return one Account with the Account number provided from the the DB.
func FindAccount(db *gorm.DB, AccountNumber string) (account Account) {
	var FoundAccount Account
	db.Where("AccountNumber=?", AccountNumber).Find(&FoundAccount)
	return FoundAccount
}

// Delete User with the NIN provided.
func Deleteuser(db *gorm.DB, nin int) (user User) {
	var userToDelete User
	db.Where("Nin=?", nin).Delete(&userToDelete)
	return userToDelete
}

// Delete Account with the Account Number provided.
func DeleteAccount(db *gorm.DB, accountNumber string) (deletedaccount Account) {
	var accountToDelete Account
	db.Where("Nin=?", accountNumber).Delete(&accountToDelete)
	return accountToDelete
}
