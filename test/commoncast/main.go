package commoncast

import (
	"time"

	"github.com/OZahed/reflectutil/test/entities"
	"github.com/OZahed/reflectutil/test/repo"
)

func EntityToRepoCardArray(en []entities.Card) ([]repo.Card, error) {
	cards := make([]repo.Card, len(en))
	for idx, eCard := range en {
		rCard, err := EntityToRepoCard(eCard)
		if err != nil {
			return nil, err
		}

		cards[idx] = rCard
	}

	return cards, nil
}

func EntityToRepoCardSlice(en []entities.Card) ([]repo.Card, error) {
	cards := make([]repo.Card, 0)
	for _, eCard := range en {
		rCard, err := EntityToRepoCard(eCard)
		if err != nil {
			return nil, err
		}

		cards = append(cards, rCard)
	}

	return cards, nil
}

func EntityToRepoCard(en entities.Card) (repo.Card, error) {
	time1 := en.Time.Format(time.RFC3339)
	inner, err := CastInnerStructs(en.Inner)
	if err != nil {
		return repo.Card{}, err
	}

	inner2, err := CastInnerStruct2(en.InnerStruct2)
	if err != nil {
		return repo.Card{}, err
	}

	return repo.Card{
		Time:         repo.NewStrTime(time1),
		Name:         en.Name,
		Age:          en.Age,
		Inner:        inner,
		InnerStruct2: inner2,
	}, nil
}

func CastInnerStructs(en entities.InnerStruct) (repo.InnerStruct, error) {
	ext := make(map[string]int)

	for k, v := range en.Extras {
		ext[k] = v
	}

	return repo.InnerStruct{
		TimeString: repo.NewStrTime(en.TimeString),
		At:         en.At,
		Name:       string(en.Name),
		Duration:   en.Duration,
		Extras:     ext,
	}, nil
}

func CastInnerStruct2(en entities.InnerStruct2) (repo.InnerStruct2, error) {
	extras := make([]repo.ExtraInfo, len(en.Extras))
	for idx, extra := range en.Extras {
		extraInfo, err := CastExtaInfo(extra)
		if err != nil {
			return repo.InnerStruct2{}, err
		}

		extras[idx] = extraInfo
	}

	return repo.InnerStruct2{
		Name:   string(en.Name),
		Extras: extras,
	}, nil
}

func CastExtaInfo(en entities.ExtraInfo) (repo.ExtraInfo, error) {
	str, err := repo.StrTimeFromEpoch(en.TimeEpoch)
	if err != nil {
		return repo.ExtraInfo{}, err
	}

	return repo.ExtraInfo{
		TimeEpoch: str,
		Slug:      en.Slug.String(),
		Writer:    en.Writer,
		Location:  en.Location,
	}, nil
}
