package iface

import (
	"errors"

	e "github.com/vault-thirteen/BitTorrentFile/models/error"
)

func InterfaceAsString(x any) (s string, err error) {
	var ok bool
	var buf []byte
	buf, ok = x.([]uint8)
	if ok {
		s = string(buf)
		return s, nil
	}

	s, ok = x.(string)
	if ok {
		return s, nil
	}

	return "", errors.New(e.ErrTypeAssertion)
}

func InterfaceAsStringArray(x any) (sa []string, err error) {
	var ok bool
	var ar []any
	ar, ok = x.([]any)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	sa = make([]string, 0, len(ar))
	var buf string
	for _, v := range ar {
		buf, err = InterfaceAsString(v)
		if err != nil {
			return nil, err
		}

		sa = append(sa, buf)
	}

	return sa, nil
}

func InterfaceAsArrayOfStringArrays(x any) (asa [][]string, err error) {
	var ok bool
	var ar []any
	ar, ok = x.([]any)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	asa = make([][]string, 0, len(ar))
	var buf []string
	for _, v := range ar {
		buf, err = InterfaceAsStringArray(v)
		if err != nil {
			return nil, err
		}

		asa = append(asa, buf)
	}

	return asa, nil
}

func InterfaceAsInt(x any) (i int, err error) {
	var ok bool
	i, ok = x.(int)
	if ok {
		return i, nil
	}

	var i64 int64
	i64, ok = x.(int64)
	if ok {
		return int(i64), nil
	}

	var i32 int32
	i32, ok = x.(int32)
	if ok {
		return int(i32), nil
	}

	var i16 int16
	i16, ok = x.(int16)
	if ok {
		return int(i16), nil
	}

	var i8 int8
	i8, ok = x.(int8)
	if ok {
		return int(i8), nil
	}

	return 0, errors.New(e.ErrTypeAssertion)
}
