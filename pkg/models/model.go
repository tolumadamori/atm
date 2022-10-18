package model

import (
	"gorm.io/gorm"
)

// User Model. Has the following fields: Nin, Name, Address, Password and a slice of Accounts of type Account. Primary field in the DB is the ID and cannot be null.
type User struct {
	gorm.Model
	Nin      int       `json:"Nin"`
	Name     string    `json:"Name"`
	Address  string    `json:"Address"`
	Password string    `json:"Password"`
	Accounts []Account `json:"Accounts" gorm:"foreignKey:AccountNumber"`
}

// Account Model. Has the following fields:  Bank, AccountType, AccountNumber and AccountBalance. Primary field in the DB is ID and cannot be null.
type Account struct {
	gorm.Model
	Bank           string  `json:"Bank"`
	AccountType    string  `json:"AccountType"`
	AccountNumber  int     `json:"AccountNumber"`
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
func FindAccount(db *gorm.DB, AccountNumber string) (account *Account) {
	var FoundAccount Account
	db.Where("account_Number=?", AccountNumber).Find(&FoundAccount)
	return &FoundAccount
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
	db.Where("AccountNumber=?", accountNumber).Delete(&accountToDelete)
	return accountToDelete
}

// Update User function. Function identifies the user to update using the Nin of the user passed to the function.
func UpdateUser(db *gorm.DB, usertoupdate *User) *User {

	db.Where("Nin=?", usertoupdate.Nin).Find(&usertoupdate).Model(&usertoupdate).Updates(map[string]interface{}{"Nin": usertoupdate.Nin, "Name": usertoupdate.Name, "Adress": usertoupdate.Address, "Password": usertoupdate.Password})
	return usertoupdate
}

// function to Update account. Account to update is identified using the Account Number.
func UpdateAccount(db *gorm.DB, accountnumber int) *Account {
	accounttoupdate := &Account{}
	db.Where("account_number=?", accountnumber).Model(&accounttoupdate).Updates(map[string]interface{}{"Bank": accounttoupdate.Bank, "AccountType": accounttoupdate.AccountType, "AccountNumber": accounttoupdate.AccountNumber})
	return accounttoupdate
}
