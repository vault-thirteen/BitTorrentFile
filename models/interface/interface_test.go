package iface

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_InterfaceAsString(t *testing.T) {
	aTest := tester.New(t)
	var x any
	var result string
	var err error

	// Test #1. Byte array.
	x = []byte("Test")
	result, err = InterfaceAsString(x)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "Test")

	// Test #2. String
	x = "Abc"
	result, err = InterfaceAsString(x)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "Abc")

	// Test #3. Unsupported type.
	x = 123
	result, err = InterfaceAsString(x)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")
}

func Test_InterfaceAsStringArray(t *testing.T) {
	aTest := tester.New(t)
	var result []string
	var err error

	// Test #1. Not an array.
	result, err = InterfaceAsStringArray(123)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []string(nil))

	// Test #2. Normal array.
	result, err = InterfaceAsStringArray([]any{"A", "B"})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{"A", "B"})

	// Test #3. Bad array.
	result, err = InterfaceAsStringArray([]any{"A", 456})
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []string(nil))
}

func Test_InterfaceAsArrayOfStringArrays(t *testing.T) {
	aTest := tester.New(t)
	var result [][]string
	var err error

	// Test #1. Not an array.
	result, err = InterfaceAsArrayOfStringArrays(123)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, [][]string(nil))

	// Test #2. Normal array.
	result, err = InterfaceAsArrayOfStringArrays([]any{[]any{"A", "B"}, []any{"C", "D"}})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, [][]string{{"A", "B"}, {"C", "D"}})

	// Test #3. Bad array.
	result, err = InterfaceAsArrayOfStringArrays([]any{[]any{"A", "B"}, 456})
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, [][]string(nil))
}

func Test_InterfaceAsInt(t *testing.T) {
	aTest := tester.New(t)
	var result int
	var err error

	// Test #1. Int.
	result, err = InterfaceAsInt(int(1))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 1)

	// Test #2. Int64.
	result, err = InterfaceAsInt(int64(2))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 2)

	// Test #3. Int32.
	result, err = InterfaceAsInt(int32(3))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 3)

	// Test #4. Int16.
	result, err = InterfaceAsInt(int16(4))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 4)

	// Test #5. Int8.
	result, err = InterfaceAsInt(int8(5))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 5)

	// Test #6. Not an int.
	result, err = InterfaceAsInt("6")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, 0)
}
