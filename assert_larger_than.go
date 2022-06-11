package test

import (
    "testing"

    "github.com/apitalist/lang"
)

// AssertLargerThan tests if the specified value is larger than the expected value.
//
// For example:
//
//     func TestAssertLargerThan(t *testing.T) {
//        actual := 3
//        expected := 2
//        test.AssertLargerThan(t, actual, expected, "Expected %d to be larger than %d", actual, expected)
//     }
func AssertLargerThan[T lang.Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
    t.Helper()
    if value <= expected {
        t.Fatalf(message, args...)
    }
}

// AssertLargerThanOrEqualTo tests if the specified value is larger than or equal to the expected value.
//
// For example:
//
//     func TestAssertLargerOrEqualThan(t *testing.T) {
//        actual := 3
//        expected := 2
//        test.AssertLargerThanOrEqualTo(t, actual, expected, "Expected %d to be larger than %d", actual, expected)
//     }
func AssertLargerThanOrEqualTo[T lang.Ordered](t *testing.T, value, expected T, message string, args ...interface{}) {
    t.Helper()
    if value < expected {
        t.Fatalf(message, args...)
    }
}
