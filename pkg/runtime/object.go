package runtime

import (
	"github.com/lambdanew/ln/pkg/common"
)

const (
	// things
	ObjectTypeNode     = 0
	ObjectTypePlace    = 1
	ObjectTypeThing    = 2
	ObjectTypeLocation = 3

	// user and "NPCs"
	ObjectTypeUser  = 10
	ObjectTypeBot   = 11
	ObjectTypeAgent = 12
	ObjectTypeTwin  = 20

	AttributeTypeString  = 1
	AttributeTypeNumber  = 2
	AttributeTypeBoolean = 3
)

// InstanceProviderFunc func() interface{}

type (
	ObjectType    int
	AttributeType int

	Operation func(...Attribute) error // FIXME this might be different e.g. not using attributes but objects ?

	Object struct {
		// names and address
		Name     string         // a 'short name', without any space
		LongName string         // a more descriptive name of the object, almost it's 'description'
		Address  common.Address // 160bit address / UUID of the object
		Type     ObjectType     // defines the kind of object

		// object hierarchy
		Parent   *Object
		Children map[string]*Object // key is the object's address

		// ownership and permissions

		// attributes
		Attributes map[string]*Attribute

		// operations
		Operations map[string]Operation
	}

	Attribute struct {
		Name  string // name of the attribute
		Type  AttributeType
		Value interface{} // some value
	}
)

// NewObject creates a minimally initialized Object
func NewObject(name string, t ObjectType, parent *Object) *Object {
	o := Object{
		Name:       name,
		LongName:   name,
		Type:       t,
		Parent:     parent,
		Children:   make(map[string]*Object),
		Attributes: make(map[string]*Attribute),
		Operations: make(map[string]Operation),
	}

	return &o
}
