package requestconfig

import "testing"

func TestCheckPathSegments(t *testing.T) {
	for path, valid := range map[string]bool{
		"a/b":           true,
		"a/..%2F..%2Fb": true,
		"a/...":         true,
		"a/b..c":        true,
		"a/%252E%252E":  true,
		"a/b?next=../c": true,
		"a/.":           false,
		"a/..":          false,
		"a/../b":        false,
		"a/..?next=b":   false,
	} {
		if err := CheckPathSegments(path); (err == nil) != valid {
			t.Errorf("CheckPathSegments(%q) = %v, want valid = %v", path, err, valid)
		}
	}
}
