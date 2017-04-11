package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

	// rsAPI := router.Group("/api/v1/resources")
	// {
	// 	rsAPI.POST("/", CreateJob)
	// 	rsAPI.GET("/", FetchAllJobs)
	// 	rsAPI.GET("/:id", FetchSingleJob)
	// 	rsAPI.PUT("/:id", UpdateJob)
	// 	rsAPI.DELETE("/:id", DeleteJob)
	// }

	router.Run(port)
}

//PredictTFImage :
func PredictTFImage(c *gin.Context) {
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
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("Predict result:%s", ret)})
}

//TestFoo :
func TestFoo(c *gin.Context) {
	ret := "bar"
	log.Println("TestFoo:", c)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ret})
}
