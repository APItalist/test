package test_test

import (
    "testing"

    "github.com/apitalist/test"
)

func TestAssertLargerThan(t *testing.T) {
    actual := 3
    expected := 2
    test.AssertLargerThan(t, actual, expected, "Expected %d to be larger than %d", actual, expected)
}

func TestAssertLargerThanOrEqualTo(t *testing.T) {
    actual := 3
    expected := 2
    test.AssertLargerThanOrEqualTo(
        t,
        actual,
        expected,
        "Expected %d to be larger than or equal to %d",
        actual,
        expected,
    )
    actual = 3
    expected = 3
    test.AssertLargerThanOrEqualTo(
        t,
        actual,
        expected,
        "Expected %d to be larger than or equal to %d",
        actual,
        expected,
    )
}
