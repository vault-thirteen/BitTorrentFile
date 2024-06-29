package cla

import (
	"errors"
	"os"
)

const (
	ErrCommandLineArgumentsCount = "invalid number of arguments"
)

// CommandLineArguments is a set of arguments received from command line.
type CommandLineArguments struct {
	ObjectType *ObjectType
	ObjectPath string
	Output     string
}

// NewCommandLineArguments is a constructor of a CommandLineArguments object.
func NewCommandLineArguments() (cla *CommandLineArguments, err error) {
	if len(os.Args) != (3 + 1) {
		return nil, errors.New(ErrCommandLineArgumentsCount)
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
