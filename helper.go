package btf

import "errors"

func interfaceAsString(x any) (s string, err error) {
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

	return "", errors.New(ErrTypeAssertion)
}

func interfaceAsStringArray(x any) (sa []string, err error) {
	var ok bool
	var ar []any
	ar, ok = x.([]any)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	sa = make([]string, 0, len(ar))
	var buf string
	for _, v := range ar {
		buf, err = interfaceAsString(v)
		if err != nil {
			return nil, err
		}

		sa = append(sa, buf)
	}

	return sa, nil
}

func interfaceAsArrayOfStringArrays(x any) (asa [][]string, err error) {
	var ok bool
	var ar []any
	ar, ok = x.([]any)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	asa = make([][]string, 0, len(ar))
	var buf []string
	for _, v := range ar {
		buf, err = interfaceAsStringArray(v)
		if err != nil {
			return nil, err
		}

		asa = append(asa, buf)
	}

	return asa, nil
}

func interfaceAsInt(x any) (i int, err error) {
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

	return 0, errors.New(ErrTypeAssertion)
}

func getSectionValueAsString(tf *BitTorrentFile, sectionName string) (sv string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return "", err
	}

	return interfaceAsString(section)
}

func getSectionValueAsStringArray(tf *BitTorrentFile, sectionName string) (sa []string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return nil, err
	}

	return interfaceAsStringArray(section)
}

func getSectionValueAsArrayOfStringArrays(tf *BitTorrentFile, sectionName string) (asa [][]string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return nil, err
	}

	return interfaceAsArrayOfStringArrays(section)
}

func getSectionValueAsInt(tf *BitTorrentFile, sectionName string) (i int, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return 0, err
	}

	return interfaceAsInt(section)
}

func removeDuplicatesFromList[T comparable](in []T) (out []T) {
	out = []T{}
	m := make(map[T]bool)

	var itemExists bool
	for _, x := range in {
		itemExists, _ = m[x]
		if !itemExists {
			m[x] = true
			out = append(out, x)
		}
	}
	return out
}
