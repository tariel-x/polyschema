package polyschema

import "encoding/json"

type TypesIdentity int8

const (
	// TypesEqual means that the type A is identical to the type B
	TypesEqual TypesIdentity = 0
	// TypesNotEqual means that the type B is neither subtype nor equal type to the A
	TypesNotEqual TypesIdentity = -1
	// TypesSubtype means that the type B is subtype of the A
	TypesSubtype TypesIdentity = 1
)

// Subtype checks if child is subtype or equal type to parent
func Subtype(parent, child JsonSchema) TypesIdentity {
	if eq := checkTypeIdentity(parent, child); eq == TypesNotEqual {
		return TypesNotEqual
	}

	switch *parent.Type {
	case Object:
		return checkObject(parent, child)
	case String:
		return checkString(parent, child)
	case Number:
		return checkNumber(parent, child)
	case Integer:
		return checkNumber(parent, child)
	case Array:
		return checkArray(parent, child)
	}

	return TypesNotEqual
}

// SubtypeRaw unmarshalls raw schemas and checks if child is subtype or equal type to parent
func SubtypeRaw(parentRaw, childRaw string) (TypesIdentity, error) {
	var parent, child JsonSchema

	if err := json.Unmarshal([]byte(parentRaw), &parent); err != nil {
		return TypesNotEqual, err
	}

	if err := json.Unmarshal([]byte(childRaw), &child); err != nil {
		return TypesNotEqual, err
	}

	return Subtype(parent, child), nil
}

func checkTypeIdentity(schema1, schema2 JsonSchema) TypesIdentity {
	if schema1.Type == nil || schema2.Type == nil {
		return TypesNotEqual
	}

	if *schema1.Type != *schema2.Type {
		return TypesNotEqual
	}

	return TypesEqual
}
