package test

import (
    "testing"

    "github.com/apitalist/lang"
)

// AssertEquals tests if the value is equal to the expected value.
//
// For example:
//
//     func TestAssertEquals(t *testing.T) {
//        a := "c"
//        b := "c"
//        test.AssertEquals(t, a, b, "%s should be equal to %s", a, b)
//     }
//
// Comparing structs:
//
// When comparing structs, the AssertEquals function behaves as expected as long as the structs are not referenced by
// pointers. For example:
//
//     func TestAssertEquals(t *testing.T) {
//        a := myTestStruct{}
//        b := myTestStruct{}
//        test.AssertEquals(t, a, b, "%v should be equal to %v", a, b)
//     }
//
// However, it won't work with pointers as the pointer addresses are not equal:
//
//     func TestAssertEquals(t *testing.T) {
//        a := &myTestStruct{}
//        b := &myTestStruct{}
//        // This will fail
//        test.AssertEquals(t, a, b, "%v should be equal to %v", a, b)
//     }
//
// Similarly, other assertions involving pointers will only succeed if the pointers point to the same address. If you
// want to compare the objects
func AssertEquals[T comparable](t *testing.T, value, expected T, message string, args ...interface{}) {
    t.Helper()
    if value != expected {
        t.Fatalf(message, args...)
    }
}

// AssertEqualsObject compares two objects using the Equals interface.
//
// Example:
//
//     type testObject struct {
//        n int
//     }
//
//     func (e *testObject) Equals(other *testObject) bool {
//        return e.n == other.n
//     }
//
//     func TestAssertEqualsObject(t *testing.T) {
//        a := &testObject{
//            1,
//        }
//        b := &testObject{
//            1,
//        }
//        test.AssertEqualsObject(t, a, b, "The two test objects should be the same.")
//     }
func AssertEqualsObject[T any, K lang.Equals[T]](
    t *testing.T,
    value K,
    expected T,
    message string,
    args ...interface{},
) {
    t.Helper()
    if !value.Equals(expected) {
        t.Fatalf(message, args...)
    }
}
