package test_test

import (
    "strings"
    "testing"

    "github.com/apitalist/test"
)

func TestAssertInSlice(t *testing.T) {
    slice := []string{"a", "b", "c"}
    element := "b"
    test.AssertInSlice(t, element, slice, "%s should be in %s", element, strings.Join(slice, ","))
}
