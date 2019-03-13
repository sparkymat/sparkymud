package main

import (
	"github.com/kr/pretty"
	"github.com/sparkymat/sparkymud/config"
)

func main() {
	appConfig := config.Load()

	pretty.Log(appConfig)
}
