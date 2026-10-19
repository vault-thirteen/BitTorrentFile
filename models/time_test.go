package models

import (
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ParseBrokenTime(t *testing.T) {
	aTest := tester.New(t)
	var result time.Time
	var err error
	var zoneName string

	{
		// Test #1A One space. Date A.
		result, err = ParseBrokenTime("1999-12-31")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)

		// Test #1B. One space. Date B.
		result, err = ParseBrokenTime("31.12.1999")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)
	}

	{
		// Test #2A. Two spaces. Date A.
		result, err = ParseBrokenTime("1999-12-31 UTC")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)

		zoneName, _ = result.Zone()
		aTest.MustBeEqual(zoneName, "UTC")

		// Test #2B. Two spaces. Date B.
		result, err = ParseBrokenTime("31.12.1999 UTC")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)

		zoneName, _ = result.Zone()
		aTest.MustBeEqual(zoneName, "UTC")

		// Test #2C. Two spaces. Date C.
		result, err = ParseBrokenTime("1999-12-31 11:22:33")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)
		aTest.MustBeEqual(result.Hour(), 11)
		aTest.MustBeEqual(result.Minute(), 22)
		aTest.MustBeEqual(result.Second(), 33)

		// Test #2D. Two spaces. Date D.
		result, err = ParseBrokenTime("31.12.1999 11:22:33")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)
		aTest.MustBeEqual(result.Hour(), 11)
		aTest.MustBeEqual(result.Minute(), 22)
		aTest.MustBeEqual(result.Second(), 33)
	}

	{
		// Test #3A. Three spaces. Date A.
		result, err = ParseBrokenTime("1999-12-31 11:22:33 UTC")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)
		aTest.MustBeEqual(result.Hour(), 11)
		aTest.MustBeEqual(result.Minute(), 22)
		aTest.MustBeEqual(result.Second(), 33)

		zoneName, _ = result.Zone()
		aTest.MustBeEqual(zoneName, "UTC")

		// Test #3B. Three spaces. Date B.
		result, err = ParseBrokenTime("31.12.1999 11:22:33 UTC")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result.Year(), 1999)
		aTest.MustBeEqual(int(result.Month()), 12)
		aTest.MustBeEqual(result.Day(), 31)
		aTest.MustBeEqual(result.Hour(), 11)
		aTest.MustBeEqual(result.Minute(), 22)
		aTest.MustBeEqual(result.Second(), 33)

		zoneName, _ = result.Zone()
		aTest.MustBeEqual(zoneName, "UTC")
	}

	{
		// Test #4. Invalid date.
		result, err = ParseBrokenTime("QWERTY")
		aTest.MustBeAnError(err)
		result, err = ParseBrokenTime("A B C D E")
		aTest.MustBeAnError(err)
	}
}
