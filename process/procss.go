package proc

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"monitor/db"
	"time"
)

func Collect(interval time.Duration) {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println()
	}

	for _, v := range processes {
		p := &db.Proc{}
		p.Pid = v.Pid
		p.Name, err = v.Name()
		p.Status, err = v.Status()
		p.CreateTime, err = v.CreateTime()
		parent, err := v.Parent()
		if err != nil {
			fmt.Printf("parent is %+v", parent)
			fmt.Println("Get parent information failed. err is", err.Error())
		}
		p.Parent = parent.Pid
		p.Uids, err = v.Uids()
		p.Gids, err = v.Gids()
		p.Tgid, err = v.Tgid()
		p.NumThreads, err = v.NumThreads()
		p.CreateTime, err = v.CreateTime()
		if err != nil {
			fmt.Println("Get process information failed. err is", err.Error())
		}

		db.Gormdb.Create(p)
	}
}
