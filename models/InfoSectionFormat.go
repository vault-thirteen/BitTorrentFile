package models

type InfoSectionFormat byte

const (
	InfoSectionFormat_Unknown    = InfoSectionFormat(0)
	InfoSectionFormat_SingleFile = InfoSectionFormat(1)
	InfoSectionFormat_MultiFile  = InfoSectionFormat(2)
)
