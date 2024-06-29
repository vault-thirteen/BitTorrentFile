package models

import (
	"net/url"
	"strings"
)

// AnnounceAddress is an announce URL.
type AnnounceAddress struct {
	URL      *url.URL
	RawData  string
	IsBroken bool
}

// NewAnnounceAddressFromString parses the string into the object of
// AnnounceAddress type.
func NewAnnounceAddressFromString(s string) (aa *AnnounceAddress, err error) {
	aa = &AnnounceAddress{
		RawData:  s,
		IsBroken: false,
	}

	aa.URL, err = url.Parse(aa.RawData)
	if err == nil {
		return aa, nil
	}

	// Some bastards place "\r" at the end of URL string.
	// Such idiots deserve to be fools forever.
	aa.URL, err = url.Parse(strings.TrimSpace(aa.RawData))
	if err == nil {
		aa.IsBroken = true
		return aa, nil
	}

	return nil, err
}

// NewAnnounceAddressListFromStringArray parses the string array into the array
// of objects of AnnounceAddress type.
func NewAnnounceAddressListFromStringArray(sa []string) (aal []AnnounceAddress, err error) {
	aal = make([]AnnounceAddress, 0, len(sa))

	var aa *AnnounceAddress
	for _, s := range sa {
		aa, err = NewAnnounceAddressFromString(s)
		if err != nil {
			return nil, err
		}

		aal = append(aal, *aa)
	}

	return aal, nil
}
