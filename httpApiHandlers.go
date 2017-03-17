package main

import (
	"errors"
	"net/http"

	"github.com/bvnk/bank/accounts"
	"github.com/bvnk/bank/appauth"
	"github.com/bvnk/bank/transactions"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
}

func getTokenFromHeader(w http.ResponseWriter, r *http.Request) (token string, err error) {
	// Get token from header
	token = r.Header.Get("X-Auth-Token")
	if token == "" {
		return "", errors.New("httpApiHandlers: Could not retrieve token from headers")
	}

	// Check token
	err = appauth.CheckToken(token)
	if err != nil {
		return "", errors.New("httpApiHandlers: Token invalid")
	}

	return
}

// Extend token
func AuthIndex(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	//Extend token
	response, err := appauth.ProcessAppAuth([]string{token, "appauth", "1"})
	Response(response, err, w, r)
	return
}

// Get token
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("User")
	password := r.FormValue("Password")

	response, err := appauth.ProcessAppAuth([]string{"0", "appauth", "2", user, password})
	Response(response, err, w, r)
	return
}

// Create auth account
func AuthCreate(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("UserIdentificationNumber")
	password := r.FormValue("Password")

	response, err := appauth.ProcessAppAuth([]string{"0", "appauth", "3", userID, password})
	Response(response, err, w, r)
	return
}

// Remove auth account
func AuthRemove(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	user := r.FormValue("User")
	password := r.FormValue("Password")

	response, err := appauth.ProcessAppAuth([]string{token, "appauth", "4", user, password})
	Response(response, err, w, r)
	return
}

func AccountIndex(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1001"})
	Response(response, err, w, r)
	return
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	// Get values from POST
	accountHolderGivenName := r.FormValue("AccountHolderGivenName")
	accountHolderFamilyName := r.FormValue("AccountHolderFamilyName")
	accountHolderDateOfBirth := r.FormValue("AccountHolderDateOfBirth")
	accountHolderIdentificationNumber := r.FormValue("AccountHolderIdentificationNumber")
	accountHolderContactNumber1 := r.FormValue("AccountHolderContactNumber1")
	accountHolderContactNumber2 := r.FormValue("AccountHolderContactNumber2")
	accountHolderEmailAddress := r.FormValue("AccountHolderEmailAddress")
	accountHolderAddressLine1 := r.FormValue("AccountHolderAddressLine1")
	accountHolderAddressLine2 := r.FormValue("AccountHolderAddressLine2")
	accountHolderAddressLine3 := r.FormValue("AccountHolderAddressLine3")
	accountHolderPostalCode := r.FormValue("AccountHolderPostalCode")
	accountType := r.FormValue("AccountType")

	req := []string{
		"0",
		"acmt",
		"1",
		accountHolderGivenName,
		accountHolderFamilyName,
		accountHolderDateOfBirth,
		accountHolderIdentificationNumber,
		accountHolderContactNumber1,
		accountHolderContactNumber2,
		accountHolderEmailAddress,
		accountHolderAddressLine1,
		accountHolderAddressLine2,
		accountHolderAddressLine3,
		accountHolderPostalCode,
		accountType,
	}

	response, err := accounts.ProcessAccount(req)
	Response(response, err, w, r)
	return
}

func AccountGet(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	vars := mux.Vars(r)
	accountId := vars["accountId"]

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1002", accountId})
	Response(response, err, w, r)
	return
}

func AccountRetrieve(w http.ResponseWriter, r *http.Request) {
	// Set these in the header as they are sensitive
	ID := r.Header.Get("X-IDNumber")
	givenName := r.Header.Get("X-GivenName")
	familyName := r.Header.Get("X-FamilyName")
	email := r.Header.Get("X-EmailAddress")

	response, err := accounts.ProcessAccount([]string{"", "acmt", "1006", ID, givenName, familyName, email})
	Response(response, err, w, r)
	return
}

func AccountGetAll(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1000"})
	Response(response, err, w, r)
	return
}

func AccountTokenPost(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	pushToken := r.FormValue("PushToken")
	platform := r.FormValue("Platform")

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1003", pushToken, platform})
	Response(response, err, w, r)
	return
}

func AccountTokenDelete(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	pushToken := r.FormValue("PushToken")
	platform := r.FormValue("Platform")

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1004", pushToken, platform})
	Response(response, err, w, r)
	return
}

func AccountSearch(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	searchTerm := r.FormValue("Search")

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1005", searchTerm})
	Response(response, err, w, r)
	return
}

func TransactionCreditInitiation(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	senderDetails := r.FormValue("SenderDetails")
	recipientDetails := r.FormValue("RecipientDetails")
	amount := r.FormValue("Amount")
	lat := r.FormValue("Lat")
	lon := r.FormValue("Lon")
	desc := r.FormValue("Desc")

	response, err := transactions.ProcessPAIN([]string{token, "pain", "1", senderDetails, recipientDetails, amount, lat, lon, desc})
	Response(response, err, w, r)
	return
}

func TransactionDepositInitiation(w http.ResponseWriter, r *http.Request) {
	basicAuthUser, basicAuthPassword, ok := r.BasicAuth()
	if !ok {
		Response("", errors.New("httpApiHandlers.TransactionDepositInitiation: Error retrieving auth headers"), w, r)
		return
	}

	if (basicAuthUser == "") || (basicAuthPassword == "") {
		Response("", errors.New("httpApiHandlers.TransactionDepositInitiation: Auth must be set"), w, r)
		return
	}

	err := appauth.CheckBasicAuth(basicAuthUser, basicAuthPassword)
	if err != nil {
		Response("", err, w, r)
		return
	}

	accountDetails := r.FormValue("AccountDetails")
	amount := r.FormValue("Amount")
	lat := r.FormValue("Lat")
	lon := r.FormValue("Lon")
	desc := r.FormValue("Desc")

	response, err := transactions.ProcessPAIN([]string{"", "pain", "1000", accountDetails, amount, lat, lon, desc, basicAuthUser, basicAuthPassword})
	Response(response, err, w, r)
	return
}

func TransactionList(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}
	// Get account number from header
	accountNumber := r.Header.Get("X-Auth-AccountNumber")
	if accountNumber == "" {
		Response("", errors.New("httpApiHandlers.TransactionList: Could not retrieve accountNumber from headers"), w, r)
		return
	}

	vars := mux.Vars(r)
	perPage := vars["perPage"]
	page := vars["page"]
	timestamp := vars["timestamp"]

	response, err := transactions.ProcessPAIN([]string{token, "pain", "1001", accountNumber, page, perPage, timestamp})
	Response(response, err, w, r)
	return
}

// Merchant accounts
// Merchant account create
func MerchantAccountCreate(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	merchantName := r.FormValue("MerchantName")
	merchantDescription := r.FormValue("MerchantDescription")
	merchantContactGivenName := r.FormValue("MerchantContactGivenName")
	merchantContactFamilyName := r.FormValue("MerchantContactFamilyName")
	merchantAddressLine1 := r.FormValue("MerchantAddressLine1")
	merchantAddressLine2 := r.FormValue("MerchantAddressLine2")
	merchantAddressLine3 := r.FormValue("MerchantAddressLine3")
	merchantCountry := r.FormValue("MerchantCountry")
	merchantPostalCode := r.FormValue("MerchantPostalCode")
	merchantBusinessSector := r.FormValue("MerchantBusinessSector")
	merchantWebsite := r.FormValue("MerchantWebsite")
	merchantContactPhone := r.FormValue("MerchantContactPhone")
	merchantContactFax := r.FormValue("MerchantContactFax")
	merchantContactEmail := r.FormValue("MerchantContactEmail")
	merchantLogo := r.FormValue("MerchantLogo")
	merchantAccountType := r.FormValue("AccountType")

	req := []string{
		token,
		"acmt",
		"1100",
		merchantName,
		merchantDescription,
		merchantContactGivenName,
		merchantContactFamilyName,
		merchantAddressLine1,
		merchantAddressLine2,
		merchantAddressLine3,
		merchantCountry,
		merchantPostalCode,
		merchantBusinessSector,
		merchantWebsite,
		merchantContactPhone,
		merchantContactFax,
		merchantContactEmail,
		merchantLogo,
		merchantAccountType,
	}

	response, err := accounts.ProcessAccount(req)
	Response(response, err, w, r)
	return
}

// Merchant account update
func MerchantAccountUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	merchantID := r.FormValue("MerchantID")
	merchantName := r.FormValue("MerchantName")
	merchantDescription := r.FormValue("MerchantDescription")
	merchantContactGivenName := r.FormValue("MerchantContactGivenName")
	merchantContactFamilyName := r.FormValue("MerchantContactFamilyName")
	merchantAddressLine1 := r.FormValue("MerchantAddressLine1")
	merchantAddressLine2 := r.FormValue("MerchantAddressLine2")
	merchantAddressLine3 := r.FormValue("MerchantAddressLine3")
	merchantCountry := r.FormValue("MerchantCountry")
	merchantPostalCode := r.FormValue("MerchantPostalCode")
	merchantBusinessSector := r.FormValue("MerchantBusinessSector")
	merchantWebsite := r.FormValue("MerchantWebsite")
	merchantContactPhone := r.FormValue("MerchantContactPhone")
	merchantContactFax := r.FormValue("MerchantContactFax")
	merchantContactEmail := r.FormValue("MerchantContactEmail")
	_ = r.FormValue("MerchantLogo")

	req := []string{
		token,
		"acmt",
		"1101",
		merchantName,
		merchantDescription,
		merchantContactGivenName,
		merchantContactFamilyName,
		merchantAddressLine1,
		merchantAddressLine2,
		merchantAddressLine3,
		merchantCountry,
		merchantPostalCode,
		merchantBusinessSector,
		merchantWebsite,
		merchantContactPhone,
		merchantContactFax,
		merchantContactEmail,
		//merchantLogo,
		"", // @FIXME We leave logo out for now, need to parse pictures nicely
		merchantID,
	}

	response, err := accounts.ProcessAccount(req)
	Response(response, err, w, r)
	return
}

// Merchat account view
func MerchantAccountView(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	vars := mux.Vars(r)
	merchantID := vars["merchantID"]

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1102", merchantID})
	Response(response, err, w, r)
	return
}

// Merchat account delete
func MerchantAccountDelete(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	vars := mux.Vars(r)
	merchantID := vars["merchantID"]
	accountID := vars["accountID"]

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1103", merchantID, accountID})
	Response(response, err, w, r)
	return
}

// Merchat account search
func MerchantAccountSearch(w http.ResponseWriter, r *http.Request) {
	token, err := getTokenFromHeader(w, r)
	if err != nil {
		Response("", err, w, r)
		return
	}

	searchTerm := r.FormValue("Search")

	response, err := accounts.ProcessAccount([]string{token, "acmt", "1104", searchTerm})
	Response(response, err, w, r)
	return
}
