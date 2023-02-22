package jobs

import (
	"api/middleware/loggo"
	"strconv"

	"github.com/robfig/cron/v3"
)

func InitJobs() {
	ObjCron := cron.New()
	Id, _ := ObjCron.AddJob("* * * * *", Test{})
	loggo.WriteLogStr("插入任务" + strconv.Itoa(int(Id)))
	ObjCron.Start()
}
