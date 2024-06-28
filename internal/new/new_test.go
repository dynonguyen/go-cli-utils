package newcli

import (
	"slices"
	"testing"
)

func TestGetDirFile(t *testing.T) {
	testCases := []struct {
		path, dir, file string
	}{
		{path: "", dir: "", file: ""},
		{path: "/", dir: "/", file: ""},
		{path: "main.go", dir: "", file: "main.go"},
		{path: "~/main.go", dir: "~/", file: "main.go"},
		{path: "dir/sub1/sub2/", dir: "dir/sub1/sub2/", file: ""},
		{path: "dir/sub1/sub2/main.go", dir: "dir/sub1/sub2/", file: "main.go"},
	}

	for _, tc := range testCases {
		if dir, file := getDirFile(tc.path); dir != tc.dir || file != tc.file {
			t.Errorf("FAIL => Input: %v, Expected: '%v', '%v' - Actual: '%v', '%v'", tc.path, tc.dir, tc.file, dir, file)
		}
	}
}

func TestGetPaths(t *testing.T) {
	testCases := []struct {
		args  []string
		paths []string
	}{
		{args: []string{}, paths: []string{}},
		{
			args: []string{
				"tmp.js",
				"tmp/",
				"tmp/main.js",
				"dir/sub1/sub2/",
				"dir/[nested1,nested2]/",
				"dir/sub1/[file.go,file2.go]",
				"dir/sub1/[file.go,file2.go,file3.go]",
			},
			paths: []string{
				"tmp.js",
				"tmp/",
				"tmp/main.js",
				"dir/sub1/sub2/",
				"dir/nested1/",
				"dir/nested2/",
				"dir/sub1/file.go",
				"dir/sub1/file2.go",
				"dir/sub1/file3.go",
			},
		},
	}

	for _, tc := range testCases {
		if paths := getPaths(&tc.args); !slices.Equal(paths, tc.paths) {
			t.Errorf("FAIL => Input: %v, Expected: '%v' - Actual: '%v'", tc.args, tc.paths, paths)
		}
	}
}
