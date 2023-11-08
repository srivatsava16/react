package funct

import (
	"awesomeProject1/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	hitCount int
	hitQueue = make([]models.Manualresponse, 8)
)

func ToQueue(hitname string, input models.Employee, response models.Response) error {
	var hitData models.Manualresponse
	hitData.Hit = hitname
	hitData.Input = input
	hitData.Manual = response
	if hitCount > 7 {
		return errors.New("hits out of bounds")
	}
	if hitCount == 7 {
		hitQueue[hitCount] = hitData
		hitCount++
		//fmt.Println(hitQueue[hitCount])
		err := ToFile(hitQueue)
		return err
	} else {
		hitQueue[hitCount] = hitData
		//fmt.Println(hitQueue[hitCount])
		hitCount++
	}
	return nil
}
func ToFile(hitQueue []models.Manualresponse) error {

	fh, err := os.OpenFile("mJson.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	log.Println("File opened")
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Println(err)
		}
	}(fh)

	jsonObj, err := json.MarshalIndent(&hitQueue, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(fh, string(jsonObj))
	if err != nil {
		return err
	}
	log.Println(string(rune(hitCount)) + " hits written to file")
	return nil
}

func Createemp(c *gin.Context) {
	var newemp models.Employee
	var response models.Response
	//img, _, _ := c.Request.FormFile("image")
	//fileByte, err := ioutil.ReadAll(img)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//str1 := string(fileByte[:])
	//if img == nil {
	//	c.JSON(http.St`atusBadRequest, "Image Required")
	//}
	//buff := make([]byte, 512)
	//_, err = img.Read(buff)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(http.DetectContentType(buff))
	//if http.DetectContentType(buff) == "image/jpeg" {
	//	c.JSON(http.StatusOK, "JPEG")
	//}

	newemp = models.Employee{
		Id:        c.PostForm("Id"),
		FirstName: c.PostForm("Firstname"),
		LastName:  c.PostForm("Lastname"),
		FullName:  c.PostForm("Fullname"),
		Phone:     c.PostForm("Phone"),
		Email:     c.PostForm("Email"),
		City:      c.PostForm("City"),
		//Img:       fileByte,
	}
	if newemp.Id == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  "Id cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Id empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//if newemp.Img == nil {
	//
	//	response = models.Response{
	//		Status: "Error",
	//		Code:   "400",
	//		Error:  "Image cannot be empty",
	//		Data:   nil,
	//	}
	//
	//	err := ToQueue("Image empty", newemp, response)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//
	//	c.JSON(http.StatusBadRequest, response)
	//	return
	//}
	if newemp.FirstName == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " FirstName cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Firstname empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if newemp.LastName == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " LastName cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Lastname empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if newemp.FullName == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " Full name cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Full name empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if newemp.Phone == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " Phone cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Phone empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if newemp.Email == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " Email cannot be empty",
			Data:   nil,
		}

		err := ToQueue("Email empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	if newemp.City == "" {

		response = models.Response{
			Status: "Error",
			Code:   "400",
			Error:  " City cannot be empty",
			Data:   nil,
		}

		err := ToQueue("City empty", newemp, response)
		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response = models.Response{
		Status: "Success",
		Code:   "200",
		Error:  "",
		Data:   "success",
	}

	err := ToQueue("Success hit", newemp, response)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, response)
	return
}

func Image(c *gin.Context) {

	body, err := c.FormFile("uploadImage")
	if err != nil {
		fmt.Println(err)

		return
	}

	ext := filepath.Ext(body.Filename)
	fmt.Println(ext)
	imgPath := "./Uploads/" + body.Filename
	if ext != ".jpeg" {
		c.JSON(http.StatusBadRequest, models.Response{
			Status: "invalid extension",
			Code:   "400",
			Error:  "Only jpeg images are allowed",
			Data:   nil,
		})
		return
	}
	err = c.SaveUploadedFile(body, imgPath)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, models.Response{
		Status: "Success",
		Code:   "400",
		Error:  "",
		Data:   body,
	})
}
