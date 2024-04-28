package repo

import (
	"io"
	"time"
)

type ExtraInfo struct {
	TimeEpoch StrTime
	Slug      string
	Writer    io.Writer
	Location  string
}

type InnerStruct struct {
	TimeString StrTime
	At         time.Time
	Extras     map[string]int
	Name       string
	Duration   time.Duration
}

type Card struct {
	Time  StrTime
	Name  string
	Inner InnerStruct
	InnerStruct2
	Age int32
}

type InnerStruct2 struct {
	Name   string
	Extras []ExtraInfo
}
