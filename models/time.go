package models

import (
	"fmt"
	"strings"
	"time"

	e "github.com/vault-thirteen/BitTorrentFile/models/error"
)

const (
	SpaceStr = " "
)

const (
	TimeFormat1A = "2006-01-02"
	TimeFormat1B = "02.01.2006"

	TimeFormat2A = "2006-01-02 MST"
	TimeFormat2B = "02.01.2006 MST"
	TimeFormat2C = "2006-01-02 15:04:05"
	TimeFormat2D = "02.01.2006 15:04:05"

	TimeFormat3A = "2006-01-02 15:04:05 MST"
	TimeFormat3B = "02.01.2006 15:04:05 MST"
)

// ParseBrokenTime tries to parse a broken time string.
func ParseBrokenTime(s string) (t time.Time, err error) {
	parts := strings.Split(s, SpaceStr)

	if len(parts) == 1 {
		t, err = time.Parse(TimeFormat1A, s)
		if err == nil {
			return t, nil
		}

		t, err = time.Parse(TimeFormat1B, s)
		if err == nil {
			return t, nil
		}
	}

	if len(parts) == 2 {
		t, err = time.Parse(TimeFormat2A, s)
		if err == nil {
			return t, nil
		}

		t, err = time.Parse(TimeFormat2B, s)
		if err == nil {
			return t, nil
		}

		t, err = time.Parse(TimeFormat2C, s)
		if err == nil {
			return t, nil
		}

		t, err = time.Parse(TimeFormat2D, s)
		if err == nil {
			return t, nil
		}
	}

	if len(parts) == 3 {
		t, err = time.Parse(TimeFormat3A, s)
		if err == nil {
			return t, nil
		}

		t, err = time.Parse(TimeFormat3B, s)
		if err == nil {
			return t, nil
		}
	}

	return t, fmt.Errorf(e.ErrFTimeFormatIsBroken, s)
}
