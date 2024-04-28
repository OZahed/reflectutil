package repo

import (
	"io"
	"time"
)

type ExtraInfo struct {
	Slug     string
	Writer   io.Writer
	Location string
}

type InnerStruct struct {
	At       time.Time
	Extras   map[string]int
	Name     string
	Duration time.Duration
}

type InnerStruct2 struct {
	Name   string
	Extras []ExtraInfo
}

type Card struct {
	Name  string
	Inner InnerStruct
	InnerStruct2
	Age int32
}
