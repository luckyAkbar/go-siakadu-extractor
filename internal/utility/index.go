package utility

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Utility struct{
	logFileTarget *os.File
}

const IMAGE_URL = "https://siakadu.unila.ac.id/uploads/fotomhs/"

var POSSIBLE_IMAGE_EXT []string = []string{".jpg", ".jpeg", ".png"}
var BASE_IMAGE_STORAGE = "./result/images/"
var LOG_FILE_LOCATION = fmt.Sprintf("./report/%s.log", time.Now().Local().String())

func New() *Utility {
	logFile, err := os.OpenFile(LOG_FILE_LOCATION, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("Failed to initialize logfile")
		log.Println(err.Error())

		os.Exit(-1)
	}

	return &Utility{
		logFileTarget: logFile,
	}
}
