package apijson

import (
	"math"
	"reflect"
	"testing"
)

func TestDecodeIntegerBounds(t *testing.T) {
	for _, tt := range []struct {
		name string
		json string
		want any
	}{
		{"signed_max", `9223372036854775807`, int64(math.MaxInt64)},
		{"signed_overflow", `9223372036854775808`, int64(math.MaxInt64)},
		{"signed_underflow", `-9223372036854775809`, int64(math.MinInt64)},
		{"unsigned_max", `18446744073709551615`, uint64(math.MaxUint64)},
		{"unsigned_overflow", `18446744073709551616`, uint64(math.MaxUint64)},
		{"unsigned_negative", `-10`, uint64(0)},
	} {
		t.Run(tt.name, func(t *testing.T) {
			for _, raw := range []string{tt.json, `"` + tt.json + `"`} {
				got := reflect.New(reflect.TypeOf(tt.want))
				if err := Unmarshal([]byte(raw), got.Interface()); err != nil {
					t.Fatalf("Unmarshal(%s): %v", raw, err)
				}
				if !reflect.DeepEqual(got.Elem().Interface(), tt.want) {
					t.Errorf("Unmarshal(%s) = %v, want %v", raw, got.Elem().Interface(), tt.want)
				}
			}
		})
	}
}
