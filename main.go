package main

import (
	"log"
	"os"
	"path/filepath"
)

//CurrentPath :
var CurrentPath string

func main() {
	//Make a forever channel which not exist for now
	var err error
	CurrentPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	port := os.Getenv("PORT")
	log.Println("-----Server Start in port=", port, " -----")
	serveHTTPAPI(":3000", forever)
}
