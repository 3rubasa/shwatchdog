package main

import (
	"fmt"
	"time"

	"github.com/3rubasa/osservices"
	"github.com/3rubasa/shwatchdog/config"
	"github.com/3rubasa/shwatchdog/watchdog"
)

const configPath = "./shagent.json"

func main() {
	var err error

	cfg, err := config.ReadFromFile(configPath)
	if err != nil {
		fmt.Println("Failed to get config: ", err)
		return
	}

	// Common
	ossvcs := osservices.NewOSServicesProvider()

	// 1 - watchdog DONE
	inetchecker := watchdog.NewInternetChecker(&cfg.Watchdog.InetChecker)
	wd := watchdog.New(&cfg.Watchdog, ossvcs, inetchecker)
	wd.Start()

	// TODO
	time.Sleep(1000 * time.Hour)
}
