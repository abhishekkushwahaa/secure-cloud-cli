package main

import (
	"github.com/abhishekkushwahaa/secure-cloud-cli/cmd"
	"github.com/abhishekkushwahaa/secure-cloud-cli/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
