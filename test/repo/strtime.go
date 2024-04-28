package repo

import (
	r "reflect"
	"time"

	ru "github.com/OZahed/reflectutil"
)

type StrTime struct {
	str string
}

func NewStrTime(s string) StrTime {
	return StrTime{s}
}

func StrTimeFromString(s string) (StrTime, error) {
	if _, err := time.Parse(time.RFC3339, s); err != nil {
		return StrTime{time.Now().Format(time.RFC3339)}, nil
	}

	return StrTime{s}, nil
}

func StrTimeFromEpoch(e int64) (StrTime, error) {
	t := time.Unix(e, 0)

	return StrTime{str: t.Format(time.RFC3339)}, nil
}

func (str StrTime) CastTo() ru.CastMap {
	t, _ := time.Parse(time.RFC3339, string(str.str))

	return ru.CastMap{
		r.TypeOf(time.Time{}): r.ValueOf(t),
		r.TypeOf(""):          r.ValueOf(string(str.str)),
		r.TypeOf(int64(0)):    r.ValueOf(t.Unix()),
	}
}

//nolint:gosimple
func (s *StrTime) ScanValue(value interface{}) error {
	var s1 StrTime
	var err error
	switch value.(type) {
	case string:
		s, _ := value.(string)
		s1, err = StrTimeFromString(s)
		if err != nil {
			return err
		}
	case int64:
		n := value.(int64)
		s1, err = StrTimeFromEpoch(n)
		if err != nil {
			return err
		}
	case time.Time:
		t := value.(time.Time)
		s1 = StrTime{t.Format(time.RFC3339)}
	}

	s.str = s1.str
	return nil
}
