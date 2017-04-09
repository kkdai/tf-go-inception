package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func serveHttpAPI(port string, existC chan bool) {
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
		jobAPI.POST("/", CreateTFImage)
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

//CreateTFImage :
func CreateTFImage(c *gin.Context) {
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// todo := Todo{Title: c.PostForm("title"), Completed: completed}
	// db, _ := Database()
	// defer db.Close()
	// db.Save(&todo)
	// c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}