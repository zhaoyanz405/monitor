package proc

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"monitor/db"
	"time"
)

func Collect() {
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
			fmt.Printf("parent is %+v, pid is %s", parent, v.Pid)
			fmt.Println("Get parent information failed. err is", err.Error())
		}
		if parent != nil {
			p.Parent = parent.Pid
		}
		p.Tgid, err = v.Tgid()
		p.NumThreads, err = v.NumThreads()
		p.CreateTime, err = v.CreateTime()
		if err != nil {
			fmt.Println("Get process information failed. err is", err.Error())
		}

		db.Gormdb.Create(p)
	}
}

func CollectByInterval(interval time.Duration) {
	for {
		go Collect()
		time.Sleep(interval)
	}
}

func TestDB() {
	p := &db.Proc{}
	p.Pid = 1
	p.Name = "test"
	p.Status = "alive"
	p.CreateTime = time.Now().UnixNano()
	p.Parent = 0
	p.Tgid = 30
	p.NumThreads = 100

	db.Gormdb.Create(p)
}
