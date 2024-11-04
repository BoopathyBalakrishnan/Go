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

//var validate *validator.Validate


func GetAllInfo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

  var info []models.Bmstbl
  //models.DB.Find(&info)
  models.DB.Find(&info)

  json.NewEncoder(w).Encode(info)
}

func CreateAccount(w http.ResponseWriter, r *http.Request){
	var input models.Bmstbl 
  
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
  
	validate = validator.New()
	err := validate.Struct(input)
  
	if err != nil {
	  utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
	  return 
	}
  
	quest, err := models.Bmstbl.CustId
	w.Header().Set("Content-Type", "application/json")
  
	json.NewEncoder(w).Encode(quest) 
  
  }

  func GetAccount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  
	id := mux.Vars(r)["custId"]
	var quest models.Bmstbl
  
	if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil{
	  utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
	  return
	}
  
	json.NewEncoder(w).Encode(quest)
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
  
	json.NewEncoder(w).Encode(quest)
  }

