package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tolumadamori/atm/pkg/config"
	model "github.com/tolumadamori/atm/pkg/models"
	"github.com/tolumadamori/atm/pkg/utils"
	"gorm.io/gorm"
)

// Connect to the DB
var db *gorm.DB = config.ConnectDB()

// Function to abstract away the process of checking for parse errors. Parses the request body into x, the first argument.
func parseErrorChecker(x interface{}, w http.ResponseWriter, r *http.Request) {
	a := utils.Parser(r, x)
	if a != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(a)
		return
	}
}

// Func to abstract away the process of writing the response. Takes our data, checks if there is an error parsing it and if there are no errors, writes it to the response.
func writer(x interface{}, w http.ResponseWriter) {
	if res, err := json.Marshal(x); err != nil {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "pkglication")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// Handlerfunc to create User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &model.User{}
	parseErrorChecker(newUser, w, r)

	createdUser := newUser.CreateUser(db)
	writer(createdUser, w)
}

// Handler func to create a new account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	newAccount := &model.Account{}
	parseErrorChecker(newAccount, w, r)

	createdAccount := newAccount.CreateAccount(db)
	writer(createdAccount, w)
}

// // Handlerfunc to deposit Money. This will work by updating the AccountBalance value of the account. The request body must contain the account number and the new account balance.
func Deposit(w http.ResponseWriter, r *http.Request) {
	accountToCredit := &model.Account{}
	parseErrorChecker(accountToCredit, w, r)

	account := model.UpdateAccount(db, accountToCredit.AccountNumber)
	writer(account, w)
}

// Handlerfunc to withdraw from the an Account balance. Pass the "Accountnumber" and the "amountToDebit" in the request body.
func Withdraw(w http.ResponseWriter, r *http.Request) {
	type input struct {
		accountNumber string
		amountToDebit string
	}
	var accountToDebit = &input{"", ""}
	parseErrorChecker(accountToDebit, w, r)
	accountNumber, amountToDebit := accountToDebit.accountNumber, accountToDebit.amountToDebit

	//find the account in the Db with the account number provided in the request body and print the current account balance.
	accountToBeDebited := model.FindAccount(db, accountNumber)
	fmt.Println(accountToBeDebited)

	currentBalancePrintout := fmt.Sprintf("this was the previous account balance %v \n", strconv.Itoa(int(accountToBeDebited.AccountBalance)))
	writer(currentBalancePrintout, w)

	//calculate the new balance that the account should have.
	convertedAmount, _ := strconv.ParseFloat(amountToDebit, 32)
	newBalance := float64(accountToBeDebited.AccountBalance) - convertedAmount

	//Update the account to be debited with the current new balance.
	accountToBeDebited.AccountBalance = float32(newBalance)
	fmt.Println(accountToBeDebited)
	//Update DB with the new Account Balance
	newAccount := accountToBeDebited

	fmt.Println(newAccount)
	account := model.UpdateAccount(db, newAccount.AccountNumber)
	newBalancePrintOut := fmt.Sprintf("This is the new Account Balance for account: %v : %v", accountNumber, strconv.Itoa(int(account.AccountBalance)))
	writer(newBalancePrintOut, w)
}

// Return the Balance of the Account number provided in the request. Request body must include account number.
func Balance(w http.ResponseWriter, r *http.Request) {
	var accountToCheck struct {
		accountNumber string
	}
	parseErrorChecker(accountToCheck, w, r)
	accountNumber := accountToCheck.accountNumber
	fmt.Println(accountToCheck)

	// find the account in the Db with the account number provided in the request body and print the current account balance.
	accountToBeChecked := model.FindAccount(db, accountNumber)
	currentAccountBalance := accountToBeChecked.AccountBalance
	fmt.Println(accountToBeChecked)

	balanceToPrint := fmt.Sprintf("The current Account balance for the account with Account Number: " + accountNumber + " is " + strconv.Itoa(int(currentAccountBalance)))
	fmt.Println(balanceToPrint)
	writer(balanceToPrint, w)
}
