package cmd

import (
	"gategaurdian/server/database/maindb"
	"gategaurdian/server/database/memorydb"

	log "github.com/sirupsen/logrus"
)

// Bootstrap function bootstraps all the needed functionality to start the server
func BootStrap() error {
	var err error

	// Initialize the required env store load the envs required for application startup
	err = memorydb.InitRequiredEnv()
	if err != nil {
		log.Error("Error loading required env : ", err)
	}

	// Initialize the memorydb required for application startup
	err = memorydb.InitMemoryDb()
	if err != nil {
		log.Error("Error initializing memorydb instance : ", err)
	}

	// Initialize the maindb required for application startup
	err = maindb.InitMainDb()
	if err != nil {
		log.Error("Error initializing memorydb instance : ", err)
	}

	return err
}
