package main

import (
	"github.com/dwburke/raid-champ-api/cmd"
	//"github.com/dwburke/raid-champ-api/db"
	"github.com/dwburke/raid-champ-api/logger"
)

func main() {
	//defer db.Close()
	defer logger.Cleanup()

	cmd.Execute()
}
