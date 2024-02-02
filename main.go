package main

import (
	"github.com/lipe7/golang-project/config"
	"github.com/lipe7/golang-project/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("go main")
	// Starting configs
	err := config.Init()
	if err != nil {
		logger.Errorf("configuration initialization failure: %v", err)
		return
	}

	// Starting router
	router.Initialize()
}
