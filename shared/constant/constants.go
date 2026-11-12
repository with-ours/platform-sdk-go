// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/with-ours/platform-sdk-go/v2/internal/encoding/json"
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

type Attribution string     // Always "attribution"
type Event string           // Always "event"
type EventProperties string // Always "event.properties"
type Funnel string          // Always "funnel"
type Insight string         // Always "insight"
type JourneyFlow string     // Always "journey-flow"
type Not string             // Always "not"
type Page string            // Always "page"
type Predicate string       // Always "predicate"

func (c Attribution) Default() Attribution         { return "attribution" }
func (c Event) Default() Event                     { return "event" }
func (c EventProperties) Default() EventProperties { return "event.properties" }
func (c Funnel) Default() Funnel                   { return "funnel" }
func (c Insight) Default() Insight                 { return "insight" }
func (c JourneyFlow) Default() JourneyFlow         { return "journey-flow" }
func (c Not) Default() Not                         { return "not" }
func (c Page) Default() Page                       { return "page" }
func (c Predicate) Default() Predicate             { return "predicate" }

func (c Attribution) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Event) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c EventProperties) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Funnel) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Insight) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c JourneyFlow) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Not) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c Page) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Predicate) MarshalJSON() ([]byte, error)       { return marshalString(c) }

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
