package helpers

import "testing"
func TestSumCollection(t *testing.T) {
	tests := []struct {
		collection []int
		sum int
		} {
			{collection: []int{1,2,3}, sum: 6 },
			{collection: []int{0,0,0}, sum: 0 },
			{collection: []int{1}, sum: 1 },
		}

	for _, tc := range tests {
		got := SumCollection(tc.collection)
		if got != tc.sum {
			t.Fatalf("Got: %d Expected: %d\n", got, tc.sum)
		}
	}


}