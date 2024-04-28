package entities

import (
	"io"
	"os"
	"time"
)

type slug struct {
	slug string
}

type Enum string

const (
	Value1 Enum = "Value1"
	Value2 Enum = "Value2"
)

type ExtraInfo struct {
	Writer   io.Writer
	Location string
}

type InnerStruct struct {
	At       time.Time
	Extras   map[string]int
	Name     Enum
	Duration time.Duration
}

type InnerStruct2 struct {
	Name   Enum
	Extras []ExtraInfo
}

type Card struct {
	Name  string
	Inner InnerStruct
	InnerStruct2
	Age int32
}

func GetEntity() Card {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		panic("could not open the file")
	}
	return Card{
		Name: "Test Value",
		Age:  1000,
		Inner: InnerStruct{
			Name:     Value2,
			At:       time.Now(),
			Duration: time.Hour,
			Extras: map[string]int{
				"key1": 1,
				"key2": 2,
			},
		},
		InnerStruct2: InnerStruct2{
			Name: Value1,
			Extras: []ExtraInfo{
				{
					Writer:   os.Stdout,
					Location: "testLoc",
				},
				{
					Writer:   f,
					Location: "testLoc",
				},
			},
		},
	}
}

func GetCards() []Card {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		panic("could not open the file")
	}
	return []Card{
		{
			Name: "Test Value",
			Age:  1000,
			Inner: InnerStruct{
				Name:     Value2,
				At:       time.Now(),
				Duration: time.Hour,
				Extras: map[string]int{
					"key1": 1,
					"key2": 2,
				},
			},
			InnerStruct2: InnerStruct2{
				Name: Value1,
				Extras: []ExtraInfo{
					{
						Writer:   os.Stdout,
						Location: "testLoc",
					},
					{
						Writer:   f,
						Location: "testLoc",
					},
				},
			},
		},
		{
			Name: "Test Value",
			Age:  1000,
			Inner: InnerStruct{
				Name:     Value2,
				At:       time.Now(),
				Duration: time.Hour,
				Extras: map[string]int{
					"key1": 1,
					"key2": 2,
				},
			},
			InnerStruct2: InnerStruct2{
				Name: Value1,
				Extras: []ExtraInfo{
					{
						Writer:   os.Stdout,
						Location: "testLoc",
					},
					{
						Writer:   f,
						Location: "testLoc",
					},
				},
			},
		},
	}
}
