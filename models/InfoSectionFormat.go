package models

// InfoSectionFormat is the format of the 'info' section.
type InfoSectionFormat byte

const (
	InfoSectionFormat_Unknown    = InfoSectionFormat(0)
	InfoSectionFormat_SingleFile = InfoSectionFormat(1)
	InfoSectionFormat_MultiFile  = InfoSectionFormat(2)
)
