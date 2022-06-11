package main

import (
	proc "monitor/process"
	"time"
)

func main() {
	proc.CollectByInterval(15 * time.Second)
}
