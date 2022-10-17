package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tolumadamori/atm/pkg/config"
	model "github.com/tolumadamori/atm/pkg/models"
	"github.com/tolumadamori/atm/pkg/utils"
	"gorm.io/gorm"
)

// Connect to the DB
var db *gorm.DB = config.ConnectDB()

// Handlerfunc to create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &model.User{}
	if err := utils.Parser(r, newUser); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(err)
		return
	}
	createdUser := newUser.CreateUser(db)
	if res, err := json.Marshal(createdUser); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Handler func to create a new account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := &model.Account{}
	if err := utils.Parser(r, newAccount); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(err)
		return
	}

	createdAccount := newAccount.CreateAccount(db)
	if res, err := json.Marshal(createdAccount); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// // Handlerfunc to deposit Money. This will work by updating the AccountBalance value of the account. The request body must contain the account number and the new account balance.
func Deposit(w http.ResponseWriter, r *http.Request) {
	var accountToCredit model.Account
	if err := utils.Parser(r, accountToCredit); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(err)
	}

	account := model.UpdateAccount(db, &accountToCredit)
	if res, err := json.Marshal(account); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Handlerfunc to withdraw from the an Account balance. Pass the "Accountnumber" and the "amountToDebit" in the request.
func Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountNumber, amountToDebit := vars["AccountNumber"], vars["ammountToDebit"]

	//find the account in the Db with the account number provided in the request body and print the current account balance.
	accountToBeDebited := model.FindAccount(db, accountNumber)
	currentAccountBalance := accountToBeDebited.AccountBalance
	balanceToPrint := []byte("this was the previous account balance" + strconv.Itoa(int(currentAccountBalance)))
	if balance, err := json.Marshal(balanceToPrint); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(balance)
	}

	//calculate the new balance that the account should have.
	convertedAmount, _ := strconv.ParseFloat(amountToDebit, 32)
	newBalance := float64(currentAccountBalance) - convertedAmount

	//Update the account to be debited with the current new balance.
	accountToBeDebited.AccountBalance = float32(newBalance)

	//Update DB with the new Account Balance

	account := model.UpdateAccount(db, accountToBeDebited)
	if res, err := json.Marshal(account); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Return the Balance of the Account number provided in the request. Request body must include account number.
func Balance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountNumber := vars["AccountNumber"]
	fmt.Println(vars)

	// find the account in the Db with the account number provided in the request body and print the current account balance.
	accountToBeChecked := model.FindAccount(db, accountNumber)
	currentAccountBalance := accountToBeChecked.AccountBalance
	fmt.Println(accountToBeChecked)
	balanceToPrint := fmt.Sprintf("The current Account balance for the account with Account Number: " + accountNumber + "is " + strconv.Itoa(int(currentAccountBalance)))
	fmt.Println(balanceToPrint)
	if balance, err := json.Marshal(balanceToPrint); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(balance)
	}

}
