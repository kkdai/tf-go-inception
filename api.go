package main

import (
	"fmt"
	"io"
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
		fooAPI.POST("/", CreatFoo)
		fooAPI.GET("/", TestFoo)
		// 	rsAPI.GET("/:id", FetchSingleJob)
		// 	rsAPI.PUT("/:id", UpdateJob)
		// 	rsAPI.DELETE("/:id", DeleteJob)
	}

	// rsAPI := router.Group("/api/v1/resources")
	// {
	// 	rsAPI.POST("/", CreateJob)
	// 	rsAPI.GET("/", FetchAllJobs)
	// 	rsAPI.GET("/:id", FetchSingleJob)
	// 	rsAPI.PUT("/:id", UpdateJob)
	// 	rsAPI.DELETE("/:id", DeleteJob)
	// }

	router.Run(":3000")
}

//PredictTFImage :
func PredictTFImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	log.Println("Receive file:", header.Filename, filename)

	byt, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Read upload file err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Error happen!"})
		return
	}

	defer file.Close()
	// fmt.Fprintf(w, "%v", header)
	f, err := os.OpenFile("./test/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	ret := TFfromForm(byt)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": fmt.Sprintf("Predict result:%s", ret)})
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// todo := Todo{Title: c.PostForm("title"), Completed: completed}
	// db, _ := Database()
	// defer db.Close()
	// db.Save(&todo)
	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

//CreatFoo :
func CreatFoo(c *gin.Context) {
	log.Println("CreatFoo:", c)
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// todo := Todo{Title: c.PostForm("title"), Completed: completed}
	// db, _ := Database()
	// defer db.Close()
	// db.Save(&todo)
	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

//TestFoo :
func TestFoo(c *gin.Context) {
	ret := "bar"
	log.Println("TestFoo:", c)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ret})
}
