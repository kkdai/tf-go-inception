package main

import (
	"flag"
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

	imagefile := flag.String("image", "", "Path of a JPEG-image to extract labels for")
	flag.Parse()
	log.Println("Path=", CurrentPath, " img=", *imagefile)

	MainTF(CurrentPath, *imagefile)

	forever := make(chan bool)
	port := os.Getenv("PORT")
	log.Println("-----Server Start in port=", port, " -----")
	serveHttpAPI(":6000", forever)
}
