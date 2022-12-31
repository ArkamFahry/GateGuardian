package main

import (
	"gategaurdian/server/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	var err error

	// Logrus log format configured to json format
	log.SetFormatter(&log.JSONFormatter{})

	err = cmd.BootStrap()
	if err != nil {
		log.Error("Error bootstrapping the server : ", err)
	}

	err = cmd.Serve()
	if err != nil {
		log.Error("Error serving the server : ", err)
	}
}
