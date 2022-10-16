package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tolumadamori/atm/pkg/config"
	model "github.com/tolumadamori/atm/pkg/models"
	"github.com/tolumadamori/atm/pkg/utils"
	"gorm.io/gorm"
)

// Connect to both DB tables
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

// // Handlerfunc to deposit funds by adding the sent value to the
// func Deposit(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	accnumber := vars["AccountNumber"]
// 	account := model.FindAccount(db2, accnumber)
// }
