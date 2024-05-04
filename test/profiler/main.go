package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/OZahed/reflectutil"
	"github.com/OZahed/reflectutil/test/commoncast"
	"github.com/OZahed/reflectutil/test/entities"
	"github.com/OZahed/reflectutil/test/repo"
	"github.com/pkg/profile"
)

const (
	FileName = "/tmp/test.txt"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	n := flag.Int("n", 1, "")
	cmd := flag.String("type", "common", "type cast mode <common | cast>")
	flag.Parse()

	fmt.Printf("running %s, %d times\n", *cmd, *n)
	if *n > 1 {
		timeIt(*cmd, *n)
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "cast":
		useCast()
	case "com", "common":
		useCommon()
	default:
		log.Fatal("not enough args (common | cast)")
	}
}

func useCast() {

	f, err := os.OpenFile(FileName, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		_ = f.Close()
		_ = os.Remove(FileName)
	}()
	card := entities.GetEntity(time.Now(), f)

	rCard := repo.Card{}

	fmt.Println("running reflectutil cast")
	err = reflectutil.TypeCast(card, &rCard)
	if err != nil {
		log.Println(err)

		return
	}
}

func useCommon() {

	f, err := os.OpenFile(FileName, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)

		return
	}

	defer func() {
		_ = f.Close()
		_ = os.Remove(FileName)
	}()
	card := entities.GetEntity(time.Now(), f)

	fmt.Println("running common cast")
	_, err = commoncast.EntityToRepoCard(card)
	if err != nil {
		log.Println(err)

		return
	}
}

func timeIt(cmd string, n int) {

	f, err := os.OpenFile(FileName, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)

		return
	}

	defer func() {
		_ = f.Close()
		_ = os.Remove(FileName)
	}()
	en := entities.GetEntity(time.Now(), f)

	if cmd == "common" || cmd == "com" {
		rCards := make([]repo.Card, n)
		for idx := range rCards {
			rCards[idx], _ = commoncast.EntityToRepoCard(en)
		}
	}

	if cmd == "cast" {
		rCards := make([]repo.Card, n)
		for idx := range rCards {
			_ = reflectutil.TypeCast(en, &rCards[idx])
		}
	}
}
