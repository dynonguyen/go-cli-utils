package trashcli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type trashCliConfig struct {
	TrashPath string
}

type UnixTime struct {
	time.Time
}

type trashInfoItem struct {
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Size      int64      `json:"size"`
	IsDir     bool       `json:"isDir"`
	CreatedAt *time.Time `json:"createdAt"`
}

var defaultConfig = trashCliConfig{TrashPath: getUserHomeDir() + "/.go-trash"}

func getUserHomeDir() string {
	if dirname, err := os.UserHomeDir(); err == nil {
		return dirname
	}

	return ""
}

func getConfig() trashCliConfig {
	return defaultConfig
}

func getTrashInfo() (map[string]trashInfoItem, error) {
	infoPath := getConfig().TrashPath + "/.trash.json"
	r := make(map[string]trashInfoItem)

	if data, err := os.ReadFile(infoPath); err != nil {
		if _, cErr := os.Create(infoPath); cErr != nil {
			return nil, fmt.Errorf("failed to create a trash info file: %v", cErr)
		}

		return r, nil
	} else {

		err := json.Unmarshal(data, &r)
		if err != nil {
			os.Remove(infoPath)
			return r, fmt.Errorf("failed to read trash info. Created a new info file %v", err)
		}
	}

	return r, nil
}

func moveFileToTrash(path string, transInfo *map[string]trashInfoItem) (*trashInfoItem, error) {
	return nil, nil
}

func TrashCli() {
	config := getConfig()
	trashPath := config.TrashPath
	verbose := false

	logError := func(err interface{}) {
		if verbose && err != nil {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(trashPath); err != nil {
		if e := os.MkdirAll(trashPath, os.ModePerm); e != nil {
			log.Fatalf("failed to create a trash folder: %v", e)
		}
	}

	trashInfo, trashInfoError := getTrashInfo()
	logError(trashInfoError)

	fmt.Println(trashInfo)
	paths := []string{"./t1.a", "./t2.a"}

	for _, path := range paths {
		moveFileToTrash(path, &trashInfo)
	}
}
