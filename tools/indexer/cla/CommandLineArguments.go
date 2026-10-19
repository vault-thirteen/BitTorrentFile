package cla

import (
	"os"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_CommandLineArgumentsCount = "invalid number of arguments"
)

// CommandLineArguments is a set of arguments received from command line.
type CommandLineArguments struct {
	ObjectType *ObjectType
	ObjectPath string
	Output     string
}

// NewCommandLineArgumentsFromOs is a constructor of a CommandLineArguments object.
func NewCommandLineArgumentsFromOs() (cla *CommandLineArguments, err error) {
	if len(os.Args) != (3 + 1) {
		return nil, errors.ErrorFromText(Err_CommandLineArgumentsCount)
	}

	cla = new(CommandLineArguments)

	cla.ObjectType, err = NewObjectType(os.Args[1])
	if err != nil {
		return nil, err
	}

	cla.ObjectPath = os.Args[2]
	cla.Output = os.Args[3]

	return cla, nil
}
