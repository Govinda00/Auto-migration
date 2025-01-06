# devtool

Created CLI tool to generate auto migration file using command with respective file.

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
# My New CLI Project
devtool/
├── cmd/
│   ├── root.go
│   ├── migrations/
│   │   ├── init.go
│   │   ├── new.go
|   |   |── migrations.go
├── internal/
│   ├── database.go (created by `devtool migrations init`)
│   └── migrations/
│       └── (migration files will be created here)
├── main.go


# Create file directory
mkdir devtool
cd devtool
git init
go mod init devtool


# load the packages

go get github.com/spf13/cobra
go get github.com/golang-migrate/migrate/v4
go get github.com/go-Create file directoryql-driver/mysql

# migration file

mkdir -p internal/migrations
touch internal/database.go

# load the cobra

cobra init --pkg-name devtool

# added migration with command
cobra add migrations
cobra add migrations init
cobra add migrations new

# docker build 
docker build -t devtool .

# run the file for auto generation file with cli

go build -o devtool
./devtool migrations init
./devtool migrations new -m "added a new column to the table"

# if you want replace ./ from command thenyou need add $path
 export PATH=$PATH:/home/govinda.kumar/devtool
 source ~/.bashrc
