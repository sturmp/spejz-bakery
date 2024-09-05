//go:build testing

package assert

import "testing"

func Same[T any](t *testing.T, collection, other []T, isSame func(T, T) bool, handleNotFound func(*testing.T, T)) {
	for _, a := range collection {
		found := false
		for _, b := range other {
			if isSame(a, b) {
				found = true
				break
			}
		}
		if !found {
			handleNotFound(t, a)
		}
	}
}
