/*
Copyright © 2025 Sidharth Chauhan
*/
package main

import (
	"github.com/sidharth-chauhan/imdb-cli/cmd"
	"github.com/sidharth-chauhan/imdb-cli/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
