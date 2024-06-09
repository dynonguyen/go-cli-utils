![Logo](https://lh3.googleusercontent.com/pw/AP1GczNLudN-Kcf1Sfd-gSOJD6DZZjiRDEiT129KD9kWeAWnsRVhyCP9yVa5HVEKzjWh2qdgIrfHweA421nOtt5pWMeUa1iqAj0rZtBJu26wNgqmrorerfFwAsHaPjGoE_ixHoZ1H308iQNRqUD21Jt6PqLq=w1440-h810-s-no-gm?authuser=0)

# Overview

Simplify the process of creating files or directories by combining the `mkdir` and `touch` commands.

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
new folder2/{{file2,file3}.go,file4.txt}
```

# Installation

```sh
go install github.com/dynonguyen/new-cli/cmd/new@latest
```

# Usage

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
