package polyscheme

import (
	"testing"
)

func int2point(x int) *int {
	return &x
}

func TestSubtypeOnlyType(t *testing.T) {
	typeName1 := String
	type1 := JsonSchema{
		Type: &typeName1,
	}
	type2 := JsonSchema{
		Type: &typeName1,
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeOnlyTypeFails(t *testing.T) {
	typeName1 := String
	typeName2 := Number
	type1 := JsonSchema{
		Type: &typeName1,
	}
	type2 := JsonSchema{
		Type: &typeName2,
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsEqual(t *testing.T) {
	typeName := Object
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a"},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeObjectsNotEqual(t *testing.T) {
	typeName := Object
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a"},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"b"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsNotEqual2(t *testing.T) {
	typeName := Object
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a", "b"},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"b"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsSubtype(t *testing.T) {
	typeName := Object
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a", "b"},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Required: []string{"a", "b", "c"},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeObjectsPropertiesEqual(t *testing.T) {
	typeName := Object
	subtypeName := String
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName,
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeObjectsPropertiesNotEqual(t *testing.T) {
	typeName := Object
	subtypeName1 := String
	subtypeName2 := Number
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName2,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesNotEqual2(t *testing.T) {
	typeName := Object
	subtypeName1 := String
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"b": JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesSubtype(t *testing.T) {
	typeName := Object
	subtypeName1 := String
	subtypeName2 := Number
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName1,
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &subtypeName1,
				},
				"b": JsonSchema{
					Type: &subtypeName2,
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeObjectsPropertiesNotEqual3(t *testing.T) {
	objType := Object
	strType := String
	numType := Number
	type1 := JsonSchema{
		Type: &objType,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &objType,
					JSTypeObj: JSTypeObj{
						Properties: map[string]JsonSchema{
							"a": JsonSchema{
								Type: &numType,
							},
						},
					},
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &objType,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &objType,
					JSTypeObj: JSTypeObj{
						Properties: map[string]JsonSchema{
							"a": JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeObjectsPropertiesSubtype2(t *testing.T) {
	objType := Object
	strType := String
	type1 := JsonSchema{
		Type: &objType,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &objType,
					JSTypeObj: JSTypeObj{
						Properties: map[string]JsonSchema{
							"a": JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &objType,
		JSTypeObj: JSTypeObj{
			Properties: map[string]JsonSchema{
				"a": JsonSchema{
					Type: &objType,
					JSTypeObj: JSTypeObj{
						Properties: map[string]JsonSchema{
							"a": JsonSchema{
								Type: &strType,
							},
							"b": JsonSchema{
								Type: &strType,
							},
						},
					},
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeStringCheck(t *testing.T) {
	typeName := String
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{
			MaxLength: int2point(1),
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{
			MaxLength: int2point(1),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeStringCheck2(t *testing.T) {
	typeName := String
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{
			MaxLength: int2point(1),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeStringCheckNotEqual(t *testing.T) {
	typeName := String
	type1 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{
			MaxLength: int2point(1),
		},
	}
	type2 := JsonSchema{
		Type: &typeName,
		JSTypeString: JSTypeString{
			MaxLength: int2point(2),
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}

func TestSubtypeArrayCheckEqual(t *testing.T) {
	arrType := Array
	strType := String
	type1 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{
					MaxLength: int2point(1),
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{
					MaxLength: int2point(1),
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesEqual)
	}
}

func TestSubtypeArrayCheckSubtype(t *testing.T) {
	arrType := Array
	strType := String
	type1 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{},
			},
		},
	}
	type2 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{
					MaxLength: int2point(1),
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesSubtype {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesSubtype)
	}
}

func TestSubtypeArrayCheckNotEqual(t *testing.T) {
	arrType := Array
	strType := String
	type1 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{
					MaxLength: int2point(2),
				},
			},
		},
	}
	type2 := JsonSchema{
		Type: &arrType,
		JSTypeArr: JSTypeArr{
			Items: &JsonSchema{
				Type: &strType,
				JSTypeString: JSTypeString{
					MaxLength: int2point(1),
				},
			},
		},
	}

	ident := Subtype(type1, type2)
	if ident != TypesNotEqual {
		t.Errorf("type1 to type2 identity is %d, expects %d", ident, TypesNotEqual)
	}
}
