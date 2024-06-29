package newcli

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

// path = 'abc/xyz/[a.go,b.go]' => ['abc/xyz/a.go', 'abc/xyz/b.go']
// path = 'abc/xyz/[sub,sub2]/' => ['abc/xyz/sub/', 'abc/xyz/sub2/']
// path = 'abc/xyz/a.go' => ['abc/xyz/a.go']
func expandPaths(path string) []string {
	if matched, _ := regexp.MatchString(`^.*\/+\[.*\]\/?$`, path); matched {
		isDir := strings.HasSuffix(path, "/")

		ss := strings.Split(path, "/")
		var last, preDir string

		if isDir {
			last = ss[len(ss)-2]
			preDir = strings.Join(ss[:len(ss)-2], "/")
		} else {
			last = ss[len(ss)-1]
			preDir = strings.Join(ss[:len(ss)-1], "/")
		}

		items := strings.Split(strings.TrimFunc(last, func(r rune) bool {
			return r == '[' || r == ']'
		}), ",")

		paths := []string{}

		for _, item := range items {
			if isDir {
				paths = append(paths, preDir+"/"+item+"/")
			} else {
				paths = append(paths, preDir+"/"+item)
			}
		}

		return paths
	}

	return []string{path}
}

func getDirFile(path string) (dir, file string) {
	if strings.HasSuffix(path, "/") {
		dir = path
		return
	}

	splits := strings.Split(path, "/")

	if len(splits) == 1 {
		file = path
		return
	}

	file = splits[len(splits)-1]
	dir = strings.Replace(path, file, "", 1)
	return
}

func newFile(path string) error {
	dir, filename := getDirFile(strings.TrimSpace(path))

	if _, err := os.Stat(path); err == nil {
		if filename != "" {
			return fmt.Errorf("file already exists: %v", path)
		} else {
			return fmt.Errorf("directory already exists: %v", path)
		}
	}

	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create a directory: %v", err)
		}
	}

	if filename != "" {
		if f, err := os.Create(path); err != nil {
			return fmt.Errorf("failed to create a file: %v", err)
		} else {
			f.Close()
		}
	}

	return nil
}

func getPaths(args *[]string) (paths []string) {
	for _, arg := range *args {
		for _, path := range expandPaths(arg) {
			if !slices.Contains(paths, path) {
				paths = append(paths, path)
			}
		}
	}

	return paths
}

func GetArgs() ([]string, bool) {
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&verbose, "v", false, "Verbose output")

	flag.Parse()
	args := flag.Args()

	return args, verbose
}

func NewCli(args []string, verbose bool) {
	for _, path := range getPaths(&args) {
		err := newFile(path)

		if verbose {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Created: " + path)
			}
		}
	}
}
