package cla

import (
	"fmt"
	"strings"
)

const (
	ObjectTypeId_File   = 1
	ObjectTypeId_Folder = 2
)

const (
	ObjectTypeName_File      = "FILE"
	ObjectTypeName_Folder    = "FOLDER"
	ObjectTypeName_Directory = "DIRECTORY"
)

const (
	ErrFObjectNameUnknown = "object name is unknown: %v"
	ErrFObjectTypeUnknown = "object type is unknown: %v"
)

// ObjectType is a type of the processed object.
// It can be either a file or a folder.
type ObjectType struct {
	id   byte
	name string
}

// NewObjectType is a constructor of the ObjectType object.
func NewObjectType(objectTypeName string) (objectType *ObjectType, err error) {
	switch strings.ToUpper(objectTypeName) {
	case ObjectTypeName_File:
		return &ObjectType{
			id:   ObjectTypeId_File,
			name: ObjectTypeName_File,
		}, nil

	case ObjectTypeName_Folder:
		return &ObjectType{
			id:   ObjectTypeId_Folder,
			name: ObjectTypeName_Folder,
		}, nil

	case ObjectTypeName_Directory:
		return &ObjectType{
			id:   ObjectTypeId_Folder,
			name: ObjectTypeName_Folder,
		}, nil

	default:
		return nil, fmt.Errorf(ErrFObjectNameUnknown, objectTypeName)
	}
}

// ID returns the numeric identifier of the object type.
func (ot *ObjectType) ID() (id byte) {
	return ot.id
}

// Name returns the textual identifier of the object type.
func (ot *ObjectType) Name() (name string) {
	return ot.name
}
