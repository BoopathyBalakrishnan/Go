package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cognizant/banking-management/models"
	"github.com/cognizant/banking-management/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/bms", GetAllInfo).Methods("GET")
	router.HandleFunc("/quest/{id}", GetAccount).Methods("GET")
	router.HandleFunc("/quest", CreateAccount).Methods("POST")
	router.HandleFunc("/quest/{id}", UpdateAccount).Methods("PUT")
	//router.HandleFunc("/quest/{id}", DeleteQuest).Methods("DELETE") 

	return router
}

func UpdateQuest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  
	id := mux.Vars(r)["id"]
	var quest models.Bmstbl
  
	if err := models.DB.Where("id = ?", custid).First(&bms).Error; err != nil{
	  utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
	  return
	}
  
	var input models.Bmstbl 
  
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
  
	validate = validator.New()
	err := validate.Struct(input)
  
	if err != nil {
	  utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
	  return 
	}
	
	models.Bmstbl.AccountType = AccountType
	models.Bmstbl.ContactNo = ContactNo
  
	models.DB.Save(&bms)
  
	json.NewEncoder(w).Encode(bms)
  }
