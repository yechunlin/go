package jobs

import (
	"github.com/robfig/cron/v3"
)

func InitJobs() {
	ObjCron := cron.New()
	ObjCron.AddJob("* * * * *", Test{})
	ObjCron.Start()
}
