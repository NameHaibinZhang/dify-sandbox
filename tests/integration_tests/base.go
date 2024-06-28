package integrationtests_test

import "testing"

func runMultipleTestings(t *testing.T, iterations int, testFunc func(*testing.T)) {
	t.Parallel()
	for i := 0; i < iterations; i++ {
		testFunc(t)
	}
}
