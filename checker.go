package polyscheme

type TypesIdentity int8

const (
	TypesEqual    TypesIdentity = 0
	TypesNotEqual TypesIdentity = -1
	TypesSubtype  TypesIdentity = 1
)

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
	case Array:
		return checkArray(parent, child)
	}

	return TypesNotEqual
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
