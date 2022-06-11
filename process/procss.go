package proc

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"time"
)

func Collect(interval time.Duration) {
	v, err := process.Processes()
	if err != nil {
		fmt.Println()
	}
	fmt.Printf("%+v", v)
}
