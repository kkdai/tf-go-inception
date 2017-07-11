package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex = &sync.Mutex{}

//serveHTTPAPI :
func serveHTTPAPI(port string, existC chan bool) {
	fmt.Println("Enter http")

	go func() {
		if err, ok := <-existC; ok {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	router := gin.Default()

	jobAPI := router.Group("/api/v1/tf-image")
	{
		jobAPI.POST("/", PredictTFImage)
	}

	fooAPI := router.Group("/api/v1/foo")
	{
		fooAPI.GET("/", TestFoo)
	}
	router.Run(port)
}

//PredictTFImage :
func PredictTFImage(c *gin.Context) {
	mutex.Lock()

	log.Println("Entry PredictTFImage..")
	file, header, err := c.Request.FormFile("upload")
	defer file.Close()

	if err != nil {
		log.Println("Parse Form failed:", err)
		return
	}
	filename := header.Filename
	log.Println("Receive file:", header.Filename, filename)

	byt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Read upload file err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error happen!"})
		return
	}
	ret := TFfromForm(byt)

	mutex.Unlock()
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("Predict result:%s", ret)})
}

//TestFoo :
func TestFoo(c *gin.Context) {
	ret := "bar"
	log.Println("TestFoo:", c)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ret})
}
