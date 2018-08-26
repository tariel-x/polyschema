package polyschema

// TypeName contains name of JSON-Schema basic type (e.g. integer, object)
type TypeName string

const (
	// JSON-Schema basic types
	Object       TypeName = "object"
	Number       TypeName = "number"
	Array        TypeName = "array"
	String       TypeName = "string"
	Unrecognized TypeName = "unrecognized"
)

// JsonSchema is type for unmarshalled json-schema
type JsonSchema struct {
	JSTypeString
	JSTypeInt
	JSTypeArr
	JSTypeObj

	Title       *string               `json:"title,omitempty"`
	Type        *TypeName             `json:"type,omitempty"`
	Enum        []string              `json:"enum,omitempty"`
	Definitions map[string]JsonSchema `json:"definitions,omitempty"`
}

// JSTypeObj is type for unmarshalled json-schema object type
type JSTypeObj struct {
	Required   []string              `json:"required,omitempty"`
	Properties map[string]JsonSchema `json:"properties,omitempty"`
}

// JSTypeString is type for unmarshalled json-schema string type
type JSTypeString struct {
	MaxLength *int    `json:"maxLength,omitempty"`
	MinLength *int    `json:"minLength,omitempty"`
	Pattern   *string `json:"pattern,omitempty"`
}

// JSTypeInt is type for unmarshalled json-schema integer type
type JSTypeInt struct {
	Minimum          *int `json:"minimum,omitempty"`
	ExclusiveMinimum *int `json:"exclusiveMinimum,omitempty"`
	Maximum          *int `json:"maximum,omitempty"`
	ExclusiveMaximum *int `json:"exclusiveMaximum,omitempty"`
	MultipleOf       *int `json:"multipleOf,omitempty"`
}

// JSTypeArr is type for unmarshalled json-schema array type
type JSTypeArr struct {
	MaxItems        *int        `json:"maxItems,omitempty"`
	MinItems        *int        `json:"minItems,omitempty"`
	Items           *JsonSchema `json:"items,omitempty"`
	Contains        *JsonSchema `json:"contains,omitempty"`
	AdditionalItems *JsonSchema `json:"additionalItems,omitempty"`
	UniqueItems     *bool       `json:"uniqueItems,omitempty"`
}
