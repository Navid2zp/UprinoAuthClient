package UprinoAuthClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Only local requests are allowed for now.
// Changes to address/server will be specified latter.
// You'll need to update the package and Init function in case of an update.
var uprinoAuthURL = "https://apis.uprino.com/auth/v1/token/check"

// args:
//      Token: Token to be checked. - Required
//      PublicKey/PrivateKey: Your api key. - Required
//      NOTE: Use PrivateKey for your backend calls.
func (t *TokenCheck) Check() (TokenCheckResponse, error) {
	responseData := TokenCheckResponse{}
	if t.PublicKey == "" && t.PrivateKey == "" {
		return responseData, errors.New("no key specified")
	} else if t.Token == "" {
		return responseData, errors.New("no token specified")
	}
	req, err := http.NewRequest("GET", uprinoAuthURL, nil)
	req.Close = true
	if err != nil {
		return responseData, err
	}

	client := &http.Client{}
	req.Header.Set("Token", t.Token)
	req.Header.Set("Public-Key", t.PublicKey)
	req.Header.Set("Private-Key", t.PrivateKey)
	res, err := client.Do(req)

	if err != nil {
		return responseData, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return responseData, err
	}
	err = json.Unmarshal(body, &responseData)
	t.Checked = true
	if err != nil {
		t.Valid = false
		return responseData, err
	}
	if responseData.Status == "ok" {
		t.Valid = true
	} else {
		t.Valid = false
	}
	return responseData, nil
}

func (t *TokenCheck) IsValid() bool {
	return t.Valid
}
