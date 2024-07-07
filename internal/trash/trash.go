package trashcli

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	utils "github.com/dynonguyen/go-cli-utils/internal"
)

// ============== Types
type trashCliConfig struct {
	TrashPath string
}

type trashInfoItem struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	IsDir     bool       `json:"isDir"`
	CreatedAt *time.Time `json:"createdAt"`
}

type trashInfoMap map[string]trashInfoItem

type trashCliOption struct {
	verbose, put, list, restore, empty, restoreAll, remove bool
}

type logErrorFunc func(err interface{})

// ============== Default config
var defaultConfig = trashCliConfig{TrashPath: getUserHomeDir() + "/.go-trash"}

const trashFileName = ".trash.json"

// ============== Utilities
func getUserHomeDir() string {
	if dirname, err := os.UserHomeDir(); err == nil {
		return dirname
	}

	return ""
}

func getConfig() trashCliConfig {
	return defaultConfig
}

func getTrashInfo(trashPath string) (trashInfoMap, error) {
	infoPath := trashPath + "/" + trashFileName
	r := make(trashInfoMap)

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

func saveTrashInfo(trashPath string, info trashInfoMap) error {
	infoPath := trashPath + "/" + trashFileName

	if jsonData, err := json.Marshal(info); err != nil {
		return err
	} else {
		if wError := os.WriteFile(infoPath, jsonData, os.ModePerm); wError != nil {
			return wError
		}
	}

	return nil
}

func getNameFromPath(path string) string {
	splits := strings.Split(strings.TrimRight(path, "/"), "/")
	return splits[len(splits)-1]
}

func generateTrashInfo(path string, currentTrash trashInfoMap) (string, *trashInfoItem, error) {
	fullPath, name := path, getNameFromPath(path)

	if fp, err := filepath.Abs(path); err == nil {
		fullPath = fp
	}

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		return name, nil, err
	}

	now := time.Now()
	info := trashInfoItem{
		Id:        utils.UniqueId(),
		Name:      name,
		Path:      fullPath,
		IsDir:     fileInfo.IsDir(),
		CreatedAt: &now,
	}

	if _, ok := currentTrash[name]; ok {
		name = fmt.Sprintf("%s_%s", info.Name, now.Format("20060102150405"))
	}

	return name, &info, nil
}

func getArgs() (opt trashCliOption, args []string) {
	flag.BoolVar(&opt.verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&opt.verbose, "v", false, "Verbose output")

	flag.BoolVar(&opt.put, "put", false, "Put files/directories to trash")
	flag.BoolVar(&opt.put, "p", false, "Put files/directories to trash")

	flag.BoolVar(&opt.list, "list", false, "List trashed files")
	flag.BoolVar(&opt.list, "l", false, "List trashed files")

	flag.BoolVar(&opt.restore, "restore", false, "Restore trashed files")

	flag.BoolVar(&opt.restoreAll, "restore-all", false, "Restore all trashed files")

	flag.BoolVar(&opt.empty, "empty", false, "Empty trash")

	flag.BoolVar(&opt.remove, "remove", false, "Permanently delete items from the trash")
	flag.BoolVar(&opt.remove, "rm", false, "Permanently delete items from the trash")

	flag.Parse()

	noFlag := true
	utils.IteratorStruct(opt, func(key string, value reflect.Value) {
		if value.Bool() && key != "verbose" {
			noFlag = false
		}
	})

	if noFlag {
		opt.put = true
	}

	return opt, flag.Args()
}

// ============== Features
func pushToTrash(paths []string, config *trashCliConfig, logError logErrorFunc) {
	trashPath := config.TrashPath
	trashInfo, trashInfoError := getTrashInfo(trashPath)
	logError(trashInfoError)

	nSuccess := 0

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, path := range paths {
		wg.Add(1)
		go func() {
			mutex.Lock()
			name, info, err := generateTrashInfo(path, trashInfo)
			mutex.Unlock()

			if err != nil {
				logError(err)
				wg.Done()
				return
			}

			if moveErr := os.Rename(info.Path, trashPath+"/"+name); moveErr != nil {
				logError(moveErr)
				wg.Done()
				return
			}

			mutex.Lock()
			trashInfo[name] = *info
			nSuccess++
			mutex.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("✅ %v file has been pushed to the trash\n", nSuccess)
	saveTrashInfo(trashPath, trashInfo)
}

// ============== Main cli
func TrashCli() {
	config := getConfig()
	option, args := getArgs()
	trashInfo, trashInfoError := getTrashInfo(config.TrashPath)

	if trashInfo == nil {
		log.Fatal("failed to get trash info: ", trashInfoError)
	}

	if _, err := os.Stat(config.TrashPath); err != nil {
		if e := os.MkdirAll(config.TrashPath, os.ModePerm); e != nil {
			log.Fatalf("failed to create a trash folder: %v", e)
		}
	}

	var logError logErrorFunc = func(err interface{}) {
		if option.verbose && err != nil {
			fmt.Println(err)
		}
	}

	var requireArgs = func() {
		if len(args) == 0 {
			log.Fatal("No paths provided")
		}
	}

	if option.put {
		requireArgs()
		pushToTrash(args, &config, logError)
		return
	}
}
