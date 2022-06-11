package test

import (
    "testing"

    "github.com/apitalist/lang"
)

// AssertInSlice searches for the search argument in the provided slice and fails if the element is not found.
//
// For example:
//
//     func TestAssertInSlice(t *testing.T) {
//        slice := []string{"a", "b", "c"}
//        element := "b"
//        test.AssertInSlice(t, element, slice, "%s should be in %s", element, strings.Join(slice, ","))
//     }
func AssertInSlice[T lang.Ordered](t *testing.T, search T, slice []T, message string, args ...interface{}) {
    t.Helper()
    for _, e := range slice {
        if e == search {
            return
        }
    }
    t.Fatalf(message, args...)
}
