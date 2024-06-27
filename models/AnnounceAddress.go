package models

import "net/url"

type AnnounceAddress struct {
	URL     *url.URL
	RawData string
}

func NewAnnounceAddressFromString(s string) (aa *AnnounceAddress, err error) {
	aa = &AnnounceAddress{
		RawData: s,
	}

	aa.URL, err = url.Parse(aa.RawData)
	if err != nil {
		return nil, err
	}

	return aa, nil
}

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
