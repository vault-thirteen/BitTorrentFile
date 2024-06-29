package models

import (
	"net/url"
	"strings"
)

type AnnounceAddress struct {
	URL     *url.URL
	RawData string
}

// NewAnnounceAddressFromString parses the string into the object of
// AnnounceAddress type.
func NewAnnounceAddressFromString(s string) (aa *AnnounceAddress, err error) {
	aa = &AnnounceAddress{
		// Some bastards place "\r" at the end of URL string.
		// Such idiots deserve to be fools forever.
		RawData: strings.TrimSpace(s),
	}

	aa.URL, err = url.Parse(aa.RawData)
	if err != nil {
		return nil, err
	}

	return aa, nil
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
