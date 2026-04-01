// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/bem-team/bem-go-sdk/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Analyze string        // Always "analyze"
type Enrich string         // Always "enrich"
type Join string           // Always "join"
type PayloadShaping string // Always "payload_shaping"
type Route string          // Always "route"
type Split string          // Always "split"
type Transform string      // Always "transform"

func (c Analyze) Default() Analyze               { return "analyze" }
func (c Enrich) Default() Enrich                 { return "enrich" }
func (c Join) Default() Join                     { return "join" }
func (c PayloadShaping) Default() PayloadShaping { return "payload_shaping" }
func (c Route) Default() Route                   { return "route" }
func (c Split) Default() Split                   { return "split" }
func (c Transform) Default() Transform           { return "transform" }

func (c Analyze) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Enrich) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Join) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c PayloadShaping) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Route) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Split) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Transform) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
