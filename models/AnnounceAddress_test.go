package models

import (
	"net/url"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewAnnounceAddressFromString(t *testing.T) {
	aTest := tester.New(t)
	var result *AnnounceAddress
	var err error

	// Test #1. Bad URL.
	// Please, note that Golang parses URLs incorrectly !!!
	result, err = NewAnnounceAddressFromString("-69-:-//kaka.xyz/path")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, (*AnnounceAddress)(nil))

	// Test #2. Valid URL.
	result, err = NewAnnounceAddressFromString("https://abc.xxx")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.RawData, "https://abc.xxx")
	aTest.MustBeEqual(result.IsBroken, false)
	aTest.MustBeDifferent(result.URL, (*url.URL)(nil))

	// Test #3. Broken URL.
	result, err = NewAnnounceAddressFromString("\rhttps://abc.xxx\r")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result.RawData, "\rhttps://abc.xxx\r")
	aTest.MustBeEqual(result.IsBroken, true)
	aTest.MustBeDifferent(result.URL, (*url.URL)(nil))
}

func Test_NewAnnounceAddressListFromStringArray(t *testing.T) {
	aTest := tester.New(t)
	var result []AnnounceAddress
	var err error

	// Test #1. Bad URL.
	// Please, note that Golang parses URLs incorrectly !!!
	result, err = NewAnnounceAddressListFromStringArray([]string{"-69-:-//kaka.xyz/path"})
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []AnnounceAddress(nil))

	// Test #2. Parseable URLs.
	result, err = NewAnnounceAddressListFromStringArray([]string{"https://abc.xxx", "\rhttps://abc.xxx\r"})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(len(result), 2)
}
