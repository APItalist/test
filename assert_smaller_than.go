package test

import (
    "testing"

    "github.com/apitalist/lang"
)

// AssertSmallerThan tests if the specified value is smaller than the expected value.
//
// For example:
//
//     func TestAssertSmallerThan(t *testing.T) {
//         actual   := 1
//         expected := 2
//         test.AssertSmallerThan(t, actual, expected, "Expected %d to be smaller than %d", actual, expected)
//     }
func AssertSmallerThan[T lang.Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
    t.Helper()
    if value >= expected {
        t.Fatalf(message, args...)
    }
}

// AssertSmallerThanOrEqualTo tests if the specified value is smaller than or equal to the expected value.
//
// For example:
//
//     func AssertSmallerThanOrEqualTo(t *testing.T) {
//         actual   := 1
//         expected := 2
//         test.AssertSmallerThanOrEqualTo(t, actual, expected, "Expected %d to be smaller than %d", actual, expected)
//     }
func AssertSmallerThanOrEqualTo[T lang.Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
    t.Helper()
    if value > expected {
        t.Fatalf(message, args...)
    }
}
