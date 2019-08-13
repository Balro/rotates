package main

import (
	"log"
	"os"
	"rotate/alert"
	"rotate/check"
	"rotate/conf"
	"rotate/status"
)

func main() {
	config := conf.ParseConfig(os.Args[1])
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	q := make(chan string)

	checker := check.NewChecker(config.Service, config.Checks)
	go checker.CronCheck(q)

	alerter := alert.NewAlerter(config.Alerts)
	go alerter.CronAlert(q)

	go status.StartHeartbeat(config.Service)

	select {}
}
