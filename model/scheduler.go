package models

import "github.com/robfig/cron/v3"

var Scheduler *cron.Cron

func init() {
	Scheduler = cron.New()
	Scheduler.Start()
}
