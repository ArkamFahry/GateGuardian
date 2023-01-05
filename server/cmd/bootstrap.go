package cmd

import (
	"gategaurdian/server/database/db"
	"gategaurdian/server/database/memorystore"

	log "github.com/sirupsen/logrus"
)

// Bootstrap function bootstraps all the needed functionality to start the server
func BootStrap() error {
	var err error

	// Initialize the required env store load the envs required for application startup
	err = memorystore.InitRequiredEnv()
	if err != nil {
		log.Error("Error loading required env : ", err)
	}

	// Initialize the memorydb required for application startup
	err = memorystore.InitMemoryDb()
	if err != nil {
		log.Error("Error initializing memorydb instance : ", err)
	}

	// Initialize the maindb required for application startup
	err = db.InitMainDb()
	if err != nil {
		log.Error("Error initializing maindb instance : ", err)
	}

	return err
}
