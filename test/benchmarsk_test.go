package test

import (
	"os"
	"testing"
	"time"

	"github.com/OZahed/reflectutil"
	"github.com/OZahed/reflectutil/test/commoncast"
	"github.com/OZahed/reflectutil/test/entities"
	"github.com/OZahed/reflectutil/test/repo"
)

func BenchmarkTypeCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TypeCast()
	}
}

func BenchmarkCommonCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CastCommon()
	}
}

func BenchmarkCommonCastArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CastCommonArray()
	}
}

func BenchmarkCastCommonSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CastCommonSliceAppend()
	}
}

func BenchmarkTypeCastHandWrittenLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CastsLoop()
	}
}

func BenchmarkTypeCastArrayReflectUtil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Casts()
	}
}

func TypeCast() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	card := entities.GetEntity(time.Now(), f)
	repoCard := repo.Card{}

	if err := reflectutil.TypeCast(card, &repoCard); err != nil {
		panic(err)
	}
}

func CastsLoop() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	cards := entities.GetCards(time.Now(), f)

	repoCards := make([]repo.Card, len(cards))

	for idx := range cards {
		if err := reflectutil.TypeCast(cards[idx], &repoCards[idx]); err != nil {
			panic(err)
		}
	}
}

func CastCommon() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	card := entities.GetEntity(time.Now(), f)

	_, err = commoncast.EntityToRepoCard(card)
	if err != nil {
		panic(err)
	}
}

func CastCommonArray() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	cards := entities.GetCards(time.Now(), f)

	_, err = commoncast.EntityToRepoCardArray(cards)
	if err != nil {
		panic(err)
	}
}

func CastCommonSliceAppend() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	cards := entities.GetCards(time.Now(), f)

	_, err = commoncast.EntityToRepoCardSlice(cards)
	if err != nil {
		panic(err)
	}
}

func Casts() {
	f, err := os.OpenFile("test.text", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = os.Remove("test.text")
	}()

	card := entities.GetCards(time.Now(), f)
	repoCard := []repo.Card{}

	if err := reflectutil.TypeCast(card, &repoCard); err != nil {
		panic(err)
	}

}
