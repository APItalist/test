package test_test

import (
    "fmt"
    "testing"

    "github.com/apitalist/test"
)

func TestAssertEquals(t *testing.T) {
    a := "c"
    b := "c"
    test.AssertEquals(t, a, b, "%s should be equal to %s", a, b)
}

func TestAssertEqualsObject(t *testing.T) {
    a := &testObject{
        1,
    }
    b := &testObject{
        1,
    }
    test.AssertEqualsObject(t, a, b, "The two test objects should be the same.")
}

type testObject struct {
    n int
}

func (e *testObject) Equals(other *testObject) bool {
    return e.n == other.n
}

func (e *testObject) String() string {
    return fmt.Sprintf("%d", e.n)
}
