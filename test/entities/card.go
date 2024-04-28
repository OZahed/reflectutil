package entities

import (
	"errors"
	"io"
	"os"
	r "reflect"
	"time"

	ru "github.com/OZahed/reflectutil"
)

type slug struct {
	slug string
}

func (s slug) String() string {
	return s.slug
}

func (s *slug) ScanValue(value interface{}) error {
	// nolint
	switch value.(type) {
	case string:
		s.slug = value.(string)
	case Enum:
		s.slug = string(value.(Enum))
	default:
		return errors.New("invalid error source type")
	}

	return nil
}

func (s slug) CastTo() ru.CastMap {
	return ru.CastMap{
		r.TypeOf(""):     r.ValueOf(s.slug),
		r.TypeOf(slug{}): r.ValueOf(s),
	}
}

type Enum string

const (
	Value1 Enum = "Value1"
	Value2 Enum = "Value2"
)

type ExtraInfo struct {
	Writer    io.Writer
	Slug      slug
	Location  string
	TimeEpoch int64
}

type InnerStruct struct {
	TimeString string
	At         time.Time
	Extras     map[string]int
	Name       Enum
	Duration   time.Duration
}

type InnerStruct2 struct {
	Name   Enum
	Extras []ExtraInfo
}

type Card struct {
	Time  time.Time
	Name  string
	Inner InnerStruct
	InnerStruct2
	Age int32
}

func GetEntity(now time.Time, f *os.File) Card {
	return Card{
		Time: now,
		Name: "Test Value",
		Age:  1000,
		Inner: InnerStruct{
			TimeString: now.Format(time.RFC3339),
			Name:       Value2,
			At:         time.Now(),
			Duration:   time.Hour,
			Extras: map[string]int{
				"key1": 1,
				"key2": 2,
			},
		},
		InnerStruct2: InnerStruct2{
			Name: Value1,
			Extras: []ExtraInfo{
				{
					Slug: slug{
						slug: "StdOut",
					},
					TimeEpoch: now.Unix(),
					Writer:    os.Stdout,
					Location:  "testLoc",
				},
				{
					Slug: slug{
						slug: "File",
					},
					TimeEpoch: now.Unix(),
					Writer:    f,
					Location:  "testLoc",
				},
			},
		},
	}
}

func GetCards(t time.Time, f *os.File) []Card {
	return []Card{
		{
			Time: t,
			Name: "Test Value",
			Age:  1000,
			Inner: InnerStruct{
				TimeString: t.Format(time.RFC3339),
				Name:       Value2,
				At:         time.Now(),
				Duration:   time.Hour,
				Extras: map[string]int{
					"key1": 1,
					"key2": 2,
				},
			},
			InnerStruct2: InnerStruct2{
				Name: Value1,
				Extras: []ExtraInfo{
					{
						Slug: slug{
							slug: "STD",
						},
						TimeEpoch: t.Unix(),
						Writer:    os.Stdout,
						Location:  "testLoc",
					},
					{
						Slug: slug{
							slug: "F",
						},
						TimeEpoch: t.Unix(),
						Writer:    f,
						Location:  "testLoc",
					},
				},
			},
		},
		{
			Time: t,
			Name: "Test Value",
			Age:  1000,
			Inner: InnerStruct{
				TimeString: t.Format(time.RFC3339),
				Name:       Value2,
				At:         time.Now(),
				Duration:   time.Hour,
				Extras: map[string]int{
					"key1": 1,
					"key2": 2,
				},
			},
			InnerStruct2: InnerStruct2{
				Name: Value1,
				Extras: []ExtraInfo{
					{
						Slug: slug{
							slug: "STDOUT",
						},
						TimeEpoch: t.Unix(),
						Writer:    os.Stdout,
						Location:  "testLoc",
					},
					{
						Slug: slug{
							slug: "FILE",
						},
						TimeEpoch: t.Unix(),
						Writer:    f,
						Location:  "testLoc",
					},
				},
			},
		},
	}
}
