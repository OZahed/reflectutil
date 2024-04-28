package test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/OZahed/reflectutil"
	"github.com/OZahed/reflectutil/test/entities"
	"github.com/OZahed/reflectutil/test/repo"
)

type Enum string

func TestTypeCastOnCard(t *testing.T) {
	f, err := os.OpenFile("tmp_test.txt", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}

		if err = os.Remove("tmp_test.txt"); err != nil {
			t.Fatal(err)
		}
	}()

	now := time.Now()
	card := entities.GetEntity(now, f)
	timeStr := repo.NewStrTime(now.Format(time.RFC3339))

	eMap := make(map[string]int)
	for k, v := range card.Inner.Extras {
		eMap[k] = v
	}

	repoCard := repo.Card{}

	want := repo.Card{
		Time: timeStr,
		Name: card.Name,
		Inner: repo.InnerStruct{
			TimeString: timeStr,
			At:         card.Inner.At,
			Extras:     eMap,
			Name:       string(card.Inner.Name),
			Duration:   card.Inner.Duration,
		},
		InnerStruct2: repo.InnerStruct2{
			Name: string(card.InnerStruct2.Name),
			Extras: []repo.ExtraInfo{
				{
					TimeEpoch: timeStr,
					Slug:      "StdOut",
					Writer:    os.Stdout,
					Location:  "testLoc",
				},
				{
					TimeEpoch: timeStr,
					Slug:      "File",
					Writer:    f,
					Location:  "testLoc",
				},
			},
		},
		Age: card.Age,
	}

	t.Run("should cast values", func(t *testing.T) {
		if err := reflectutil.TypeCast(card, &repoCard); err != nil {
			t.Errorf("TypeCast() = %s, wanted nil", err.Error())
		}

		if !reflect.DeepEqual(repoCard, want) {
			t.Errorf("\n\nwant: %#v,\n\n got: %#v", want, repoCard)
		}
	})

}

func TestTypeCastToEntityCard(t *testing.T) {
	f, err := os.OpenFile("tmp_test.txt", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}

		if err = os.Remove("tmp_test.txt"); err != nil {
			t.Fatal(err)
		}
	}()

	now, err := time.Parse(time.RFC3339, "2024-04-28T15:32:10+03:30")
	if err != nil {
		t.Fatal(err)
	}
	timeStr := repo.NewStrTime(now.Format(time.RFC3339))
	card := entities.GetEntity(now, f)

	eMap := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	eCard := entities.Card{}

	reoCard := repo.Card{
		Time: timeStr,
		Name: card.Name,
		Inner: repo.InnerStruct{
			TimeString: timeStr,
			At:         card.Inner.At,
			Extras:     eMap,
			Name:       string(card.Inner.Name),
			Duration:   card.Inner.Duration,
		},
		InnerStruct2: repo.InnerStruct2{
			Name: string(card.InnerStruct2.Name),
			Extras: []repo.ExtraInfo{
				{
					TimeEpoch: timeStr,
					Slug:      "StdOut",
					Writer:    os.Stdout,
					Location:  "testLoc",
				},
				{
					TimeEpoch: timeStr,
					Slug:      "File",
					Writer:    f,
					Location:  "testLoc",
				},
			},
		},
		Age: card.Age,
	}

	want := entities.Card{
		Time: now,
		Name: "Test Value",
		Age:  1000,
		Inner: entities.InnerStruct{
			TimeString: now.Format(time.RFC3339),
			Name:       card.Inner.Name,
			At:         card.Inner.At,
			Duration:   time.Hour,
			Extras:     eMap,
		},
		InnerStruct2: entities.InnerStruct2{
			Name: card.InnerStruct2.Name,
			Extras: []entities.ExtraInfo{
				{
					Slug:      card.InnerStruct2.Extras[0].Slug,
					TimeEpoch: now.Unix(),
					Writer:    os.Stdout,
					Location:  "testLoc",
				},
				{
					Slug:      card.InnerStruct2.Extras[1].Slug,
					TimeEpoch: now.Unix(),
					Writer:    f,
					Location:  "testLoc",
				},
			},
		},
	}

	t.Run("should cast values", func(t *testing.T) {
		if err := reflectutil.TypeCast(reoCard, &eCard); err != nil {
			t.Errorf("TypeCast() = %s, wanted nil", err.Error())
		}

		if !reflect.DeepEqual(eCard, want) {
			t.Errorf("\n\nwant: %#v,\n\n got: %#v", want, eCard)
		}
	})

}
