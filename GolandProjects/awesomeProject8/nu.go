package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	Url = "https://savesolution-stag.vishwamcorp.com/v1/string_match"
)

type Input struct {
	AppType       string `json:"app_type"`
	AppId         string `json:"app_id"`
	AccountNumber string `json:"account_number"`
	Ifsc          string `json:"ifsc"`
}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Error      string `json:"error"`
	Data       Data   `json:"data"`
}
type Data struct {
	AccountExists   bool    `json:"account_exists"`
	AmountDeposited float64 `json:"amount_deposited"`
	Message         string  `json:"message"`
	NameAtBank      string  `json:"name_at_bank"`
	ReferenceId     string  `json:"reference_id"`
}

type Manualresponse struct {
	HitName        string
	Input          Input
	ManualResponse Response `json:"response"`
}

type Externalresponse struct {
	HitName          string
	Input            Input
	ManualResponse   Response
	ExternalResponse Response
	HitMatch         bool
}

func main() {

	var Hitdata []Manualresponse
	//read the file
	fd, err := os.ReadFile("testhit.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(fd, &Hitdata)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(Hitdata)
	var extHitData = make([]Externalresponse, len(Hitdata))
	fmt.Println("Starting to send request")
	for i, val := range Hitdata {

		inp := val.Input
		client := http.Client{}
		body := &bytes.Buffer{}

		writer := multipart.NewWriter(body)
		fmt.Println("setting post-data")

		err = writer.WriteField("app_type", inp.AppType)
		if err != nil {
			log.Println(err)
		}

		err = writer.WriteField("app_id", inp.AppId)
		if err != nil {
			log.Println(err)

		}
		err = writer.WriteField("account_number", inp.AccountNumber)
		if err != nil {
			log.Println(err)
		}
		err = writer.WriteField("ifsc", inp.Ifsc)
		if err != nil {
			log.Println(err)
		}

		err = writer.Close()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Hitting Api with " + val.HitName)
		req, err := http.NewRequest("POST", Url, body)
		if err != nil {
			log.Println(err)
		}
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJBMTI4S1ciLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.mqP96pUu0LBMSDYtTrMjtl57FvScGUs5w1fctuI2Bl7ToiKNEHZz0Q.iAcVhc1faymPVA60Yh17wQ.1b-4FMEb61yjfhL-rkUhVJrigtUGNtGvUuyC_h6FiT-FQXu7lfV1gRjlxVQygAXHrwM2xETfZhnusNoEJTqwy4cY7yJVQwMPCkDoKRP0QjY.Liz9l8lS4-_69KUC5dzf1g")
		req.Header.Set("Content-Type", writer.FormDataContentType())
		extResp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer extResp.Body.Close()
		//fmt.Println(extResp.Body)
		bodyBytes, err := io.ReadAll(extResp.Body)
		if err != nil {
			log.Println(err)
		}
		var extRespData Response
		err = json.Unmarshal(bodyBytes, &extRespData)
		if err != nil {
			log.Println("unmarshall - ", err)
		}

		if extRespData.Status == val.ManualResponse.Status && extRespData.StatusCode == val.ManualResponse.StatusCode && extRespData.Error == val.ManualResponse.Error && extRespData.Data == val.ManualResponse.Data {
			extHitData[i].HitMatch = true
		} else {
			extHitData[i].HitMatch = false
		}

		extHitData[i].HitName = val.HitName
		extHitData[i].Input = val.Input
		extHitData[i].ManualResponse = val.ManualResponse
		extHitData[i].ExternalResponse = extRespData

	}

	fh, err := os.OpenFile("externalJson.json", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	md, err := json.MarshalIndent(&extHitData, "", "  ")
	_, err = fmt.Fprintln(fh, string(md))
	if err != nil {
		log.Println(err)
	}
}
