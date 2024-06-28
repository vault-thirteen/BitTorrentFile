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

type ObjectType struct {
	id   byte
	name string
}

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

func (ot *ObjectType) ID() (id byte) {
	return ot.id
}

func (ot *ObjectType) Name() (name string) {
	return ot.name
}
