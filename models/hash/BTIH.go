package hash

import (
	"github.com/vault-thirteen/auxie/hash"
)

// BtihData is the BitTorrent Info Hash (BTIH) check sum stored both as a text
// and as an array of bytes. This is an original first BTIH.
type BtihData struct {
	Bytes hash.Sha1Sum
	Text  string
}

// BtihData2 is the BitTorrent Info Hash (BTIH) check sum stored both as a text
// and as an array of bytes. This is the second version of BTIH.
type BtihData2 struct {
	Bytes hash.Sha256Sum
	Text  string
}
