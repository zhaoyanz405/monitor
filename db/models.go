package db

import "gorm.io/gorm"

type Proc struct {
	gorm.Model
	Pid        int32
	Name       string
	Status     string
	Parent     int32
	Tgid       int32
	NumThreads int32
	CreateTime int64
}
