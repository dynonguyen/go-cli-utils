![Logo](https://www.dropbox.com/scl/fi/3jf52ewmmzejc6cneum81/go-cli.jpeg?rlkey=54pm45ym5el4wxp228uyt46bn&st=jyxmy80l&raw=1)

**Some cli utilities are written in go.**

# New Cli

`new` - Simplify the process of creating files or directories by combining the `mkdir` and `touch` commands.

🥹 Using `mkdir` & `touch`

```sh
# Netesd folder
mkdir -p folder/sub1/sub2
touch folder/sub1/sub2/file.go

# File
touch file2.go

# directory
mkdir dir
mkdir -p dir/sub1

# Multiple files
mkdir -p folder2
cd folder2
touch file2.go file3.go file4.txt
```

☕ Using `new`

```sh
# Netesd folder
new folder/sub1/sub2/file.go

# File
new file2.go

# directory
new dir/
new dir/sub1/

# Multiple files & directories
new folder2/[file2.go,file3.go,file4.txt]
```

### Installation

```sh
go install github.com/dynonguyen/go-cli-utils/cmd/new@latest
```

### Usage

```sh
# Create file
new file.go
new file.go file2.go file3.js

# Create directory (end with /)
new dir/
new dir/sub1/sub2/ dir2/
new "dir/[sub1,sub2]/" # Equivalent: new dir/sub1/ dir/sub2/

# Create file in directory
new dir/sub1/file.go
new "dir/sub1/[file.go,file2.go]" # Equivalent: new dir/sub1/file1.go dir/sub2/file2.go

# Space character (surrounded by double quotes)
new "dir/orange cat/cat.go"

# Combination
new "dir/[file.go,file2.go]" dir/sub1/file.go file3.js
```

# Trash cli

`trash` is a simple, fast, much safer replacement of bash `rm`.
