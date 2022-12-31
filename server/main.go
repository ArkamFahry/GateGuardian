package main

import (
	"gategaurdian/server/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	var err error

	// Logrus log format configured to json format
	log.SetFormatter(&log.JSONFormatter{})

	// Bootstrap function bootstraps all the needed functionality to start the server
	err = cmd.BootStrap()
	if err != nil {
		log.Error("Error bootstrapping the server : ", err)
	}

	// Serve function servers the newly bootstrapped server
	err = cmd.Serve()
	if err != nil {
		log.Error("Error serving the server : ", err)
	}
}
