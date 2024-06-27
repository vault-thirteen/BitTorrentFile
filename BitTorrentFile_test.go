package btf

//TODO

/*
func Test_DecodedObject_CalculateBtih(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var object DecodedObject
	var expectedBtihText string
	var expectedBtihBytes Sha1Sum

	// Test #1.
	{
		object = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("info"),
					Value: "Just a Test.",
				},
			},
		}
		expectedBtihText = "6f1ef4ba8a877d657378dbbb78badfd2eaacf2a2"
		var ba []byte
		// "Just a Test." -> "12:Just a Test." -> (SHA-1)
		ba, err = hex.DecodeString(expectedBtihText)
		aTest.MustBeNoError(err)
		copy(expectedBtihBytes[:], ba)
		err = object.CalculateBtih()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(object.BTIH.Text, expectedBtihText)
		aTest.MustBeEqual(object.BTIH.Bytes, expectedBtihBytes)
	}

	// Test #2.
	{
		object = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("no-info"),
					Value: "",
				},
			},
		}
		err = object.CalculateBtih()
		aTest.MustBeAnError(err)
	}
}

func Test_DecodedObject_GetSection(t *testing.T) {
	var aTest = tester.New(t)
	var output any
	var err error
	var input DecodedObject
	var outputExpected any

	// Test #1. Positive.
	{
		input = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("aaa"),
					Value: nil,
				},
				{
					Key:   []byte("bbb"),
					Value: 123,
				},
				{
					Key:   []byte("INFO"),
					Value: uint8(255),
				},
				{
					Key:   []byte("info"),
					Value: int16(101),
				},
				{
					Key:   []byte("section_name"),
					Value: "John Lennon",
				},
			},
		}
		outputExpected = "John Lennon"
		output, err = input.GetSection("section_name")
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(output, outputExpected)
	}

	// Test #2. Section is absent.
	{
		input = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("zzz"),
					Value: nil,
				},
			},
		}
		output, err = input.GetInfoSection()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(output, nil)
	}

	// Test #3. Type cast error.
	{
		input = DecodedObject{
			RawObject: time.Time{},
		}
		output, err = input.GetInfoSection()
		aTest.MustBeAnError(err)
	}
}

func Test_DecodedObject_GetInfoSection(t *testing.T) {
	var aTest = tester.New(t)
	var output any
	var err error
	var input DecodedObject
	var outputExpected any

	// Test #1. Positive.
	{
		input = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("aaa"),
					Value: nil,
				},
				{
					Key:   []byte("bbb"),
					Value: 123,
				},
				{
					Key:   []byte("INFO"),
					Value: uint8(255),
				},
				{
					Key:   []byte("info"),
					Value: int16(101),
				},
				{
					Key:   []byte("ccc"),
					Value: "John",
				},
			},
		}
		outputExpected = int16(101)
		output, err = input.GetInfoSection()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(output, outputExpected)
	}

	// Test #2. Section is absent.
	{
		input = DecodedObject{
			RawObject: []DictionaryItem{
				{
					Key:   []byte("zzz"),
					Value: nil,
				},
			},
		}
		output, err = input.GetInfoSection()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(output, nil)
	}

	// Test #3. Type cast error.
	{
		input = DecodedObject{
			RawObject: time.Time{},
		}
		output, err = input.GetInfoSection()
		aTest.MustBeAnError(err)
	}
}
*/
