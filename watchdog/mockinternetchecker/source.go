package mockinternetchecker

//go:generate mockgen -destination=./mockinternetchecker.go -package=mockinternetchecker github.com/3rubasa/shwatchdog/watchdog InternetChecker
