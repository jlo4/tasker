package main

import (
	tasker "github.com/jlo4/tasker/cmd"
	dbHandler "github.com/jlo4/tasker/cmd/db"
)

func main() {
	dbHandler.Connect()
	tasker.Show()
}
