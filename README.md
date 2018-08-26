# Polyschema

JSON-Schema based type resolver for golang.

Type here is any standard json-schema. However standard specification do not consider 
any other cases except data validation via schema. This library provides realization of 
subtyping and type-checking concept for schemas.

Subtype `B` of type `A` here is the new schema describing document `b` which can be 
successfully valided with both `A` and `B`. Though document `a` described by `A` is not 
committed to be valid for schema `B`. This is called structural type system.

### Example

**Type A**

```json
{ "type": "string" }
```

**Type B**

```json
{ "type": "string", "maxLength": 10 }
```

**Type C**

```json
{ "type": "integer" }
```

Type `B` is the subtype of `A`: `A :> B`. 
But `C` is neither subtype not equal type of `A`.

### Complex example

**Type A**

```json
{
  "title": "Person",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "required": [
    "name"
  ]
}
```

**Type B**

```json
{
  "title": "Student",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    },
    "course": {
      "type": "integer"
    }
  },
  "required": [
    "name",
    "course"
  ]
}
```

**Type C**

```json
{
  "title": "Person_not_strict",
  "type": "object",
  "properties": {
    "name": {
      "type": "string"
    }
  }
}
```

Type `B` is the subtype of `A`: `A :> B`. 
But `C` is neither subtype not equal type of `A`.

## Usage of the package

Install

`go get -u github.com/tariel-x/polyschema`

Use

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tariel-x/polyschema"
)

func main() {
	typeA := polyschema.JsonSchema{}
	typeB := polyschema.JsonSchema{}
	typeC := polyschema.JsonSchema{}

	json.Unmarshal([]byte(`{ "type": "string" }`), &typeA)
	json.Unmarshal([]byte(`{ "type": "string", "maxLength": 10 }`), &typeB)
	json.Unmarshal([]byte(`{ "type": "integer" }`), &typeC)

	fmt.Println("A :> B ", polyschema.Subtype(typeA, typeB)) // 1, that means subtype
	fmt.Println("A :> C ", polyschema.Subtype(typeA, typeC)) // -1, that means not subtype
	fmt.Println("A = A ", polyschema.Subtype(typeA, typeA)) // 0, that means identity
}
```