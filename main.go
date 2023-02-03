package main

import (
	"fmt"
	"time"

	"github.com/3rubasa/osservices"
	"github.com/3rubasa/shwatchdog/config"
)

const configPath = "./shagent.json"

func main() {
	var err error

	_, err = config.ReadFromFile(configPath)
	if err != nil {
		fmt.Println("Failed to get config: ", err)
		return
	}

	// Common
	_ = osservices.NewOSServicesProvider()

	// 1 - watchdog DONE
	// inetchecker := watchdog.NewInternetChecker(&cfg.Watchdog.InetChecker)
	// wd := watchdog.New(&cfg.Watchdog, osservices, inetchecker)
	// wd.Start()

	// TODO
	time.Sleep(1000 * time.Hour)
}
