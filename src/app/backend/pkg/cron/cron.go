package cron

import (
	"time"

	"github.com/robfig/cron"
	"github.com/ustack/Yunus/src/app/backend/models"
	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
)

var c = cron.New()

func NewCron() {
	var (
		entry *cron.Entry
		err   error
	)

	if setting.Cron.UpdateMirror.Enabled {
		entry, err = c.AddFunc(setting.Cron.UpdateMirror.Schedule, models.GeneratorconnStr
		if setting.Cron.UpdateMirror.RunAtStart {
			entry.Prev = time.Now()
			go models.GeneratorconnStr()
		}
	}

  c.Start()
}

// ListTasks returns all running cron tasks.
func ListTasks() []*cron.Entry {
	return c.Entries()
}
