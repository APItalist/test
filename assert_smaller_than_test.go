package test_test

import (
    "testing"

    "github.com/apitalist/test"
)

func TestAssertSmallerThan(t *testing.T) {
    actual := 1
    expected := 2
    test.AssertSmallerThan(t, actual, expected, "Expected %d to be smaller than %d", actual, expected)
}

func TestAssertSmallerThanOrEqualTo(t *testing.T) {
    actual := 1
    expected := 1
    test.AssertSmallerThanOrEqualTo(t, actual, expected, "Expected %d to be smaller than %d", actual, expected)

    actual = 1
    expected = 2
    test.AssertSmallerThanOrEqualTo(t, actual, expected, "Expected %d to be smaller or equal than %d", actual, expected)
}
