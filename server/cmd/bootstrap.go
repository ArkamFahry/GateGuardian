package cmd

import (
	"gategaurdian/server/database/memorydb"
	"gategaurdian/server/env"

	log "github.com/sirupsen/logrus"
)

func BootStrap() error {
	var err error

	// Initialize the required env store load the envs required for application startup
	err = env.InitRequiredEnv()
	if err != nil {
		log.Error("Error loading required env : ", err)
	}

	err = memorydb.InitMemStore()
	if err != nil {
		log.Error("Error initializing memorydb instance : ", err)
	}

	return err
}
