package trashcli

import (
	"fmt"
	"log"
	"os"
)

type TrashCliConfig struct {
	TrashPath string
}

func getUserHomeDir() string {
	if dirname, err := os.UserHomeDir(); err == nil {
		return dirname
	}

	return ""
}

var defaultConfig = TrashCliConfig{TrashPath: getUserHomeDir() + "/.go-trash"}

func getConfig() TrashCliConfig {
	return defaultConfig
}

func TrashCli() {
	config := getConfig()
	trashPath := config.TrashPath

	fmt.Println(trashPath)

	if _, err := os.Stat(trashPath); err != nil {
		if e := os.MkdirAll(trashPath, os.ModePerm); e != nil {
			log.Fatalf("failed to create a trash folder: %v", e)
		}
	}

	fmt.Println("RUN")
}
