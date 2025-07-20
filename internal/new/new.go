package newcli

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type NewCliFlag struct {
	Verbose bool
}

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

		paths := make([]string, 0, len(items))

		for _, item := range items {
			path := preDir + "/" + item

			if isDir {
				path += "/"
			}

			paths = append(paths, path)
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
		}

		return fmt.Errorf("directory already exists: %v", path)
	}

	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create a directory: %v", err)
		}
	}

	if filename != "" {
		f, err := os.Create(path)

		if err != nil {
			return fmt.Errorf("failed to create a file: %v", err)
		}

		f.Close()
	}

	return nil
}

func getPaths(args []string) []string {
	paths := make([]string, 0, len(args))

	for _, arg := range args {
		for _, path := range expandPaths(arg) {
			if !slices.Contains(paths, path) {
				paths = append(paths, path)
			}
		}
	}

	return paths
}

func parseFlags() *NewCliFlag {
	flags := &NewCliFlag{}

	flag.BoolVar(&flags.Verbose, "verbose", false, "Verbose output")
	flag.BoolVar(&flags.Verbose, "v", false, "Verbose output")

	flag.Parse()

	return flags
}

func parseArgs() []string {
	flag.Parse()
	return flag.Args()
}

func Execute() {
	flags := parseFlags()
	args := parseArgs()

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No paths provided\n")
		os.Exit(1)
	}

	for _, path := range getPaths(args) {
		err := newFile(path)
		if flags.Verbose {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Created: " + path)
			}
		}
	}
}
