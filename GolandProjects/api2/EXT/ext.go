package main

import (
	models "awesomeProject4/struct"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	Host     = "http://localhost"
	Port     = ":8080"
	Endpoint = "/newstuu"
	Url      = Host + Port + Endpoint
)

func init() {
	file, err := os.OpenFile("errLog.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	var manHitData []models.Manualresponse

	fd, err := os.ReadFile("/home/sukshi14/GolandProjects/api2/main/api.json")
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(fd, &manHitData)
	if err != nil {
		log.Println(err)
	}
	var extHitData = make([]models.ExtJson, len(manHitData))

	for i, val := range manHitData {

		inp := val.Input
		client := http.Client{Timeout: 3 * time.Second}
		body := &bytes.Buffer{}

		writer := multipart.NewWriter(body)
		fmt.Println("setting post-data")

		err := writer.WriteField("Id", inp.Id)
		if err != nil {
			log.Println(err)
		}

		err = writer.WriteField("Firstname", inp.Firstname)
		if err != nil {
			log.Println(err)

		}
		err = writer.WriteField("Lastname", inp.Lastname)
		if err != nil {
			log.Println(err)
		}
		err = writer.WriteField("Fullname", inp.Username)
		if err != nil {
			log.Println(err)
		}
		err = writer.WriteField("Phone", inp.Password)
		if err != nil {
			log.Println(err)
		}
		err = writer.WriteField("Email", inp.Address)
		if err != nil {
			log.Println(err)
		}
		err = writer.WriteField("City", inp.Place)
		if err != nil {
			log.Println(err)
		}

		err = writer.Close()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Hitting Api with " + val.Hit)
		req, err := http.NewRequest("POST", Url, bytes.NewBuffer(body.Bytes()))
		if err != nil {
			log.Println(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		extResp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer extResp.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(extResp.Body)
		bodyStr := string(bodyBytes)
		fmt.Println("Resp body: " + bodyStr)

		var extRespData models.Response
		err = json.Unmarshal(bodyBytes, &extRespData)
		if err != nil {
			log.Println(err)
		}

		if extRespData.Status == val.Manual.Status && extRespData.Code == val.Manual.Code && extRespData.Error == val.Manual.Error && extRespData.Data == val.Manual.Data {
			extHitData[i].HitMatch = true
		} else {
			extHitData[i].HitMatch = false
		}

		extHitData[i].HitName = val.Hit
		extHitData[i].Input = val.Input
		extHitData[i].ManualResponse = val.Manual
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
