package fuelcounterupper

import "testing"

func TestCalculateFuelByMass(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{mass: 12, fuel: 2},
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 654},
		{mass: 100756, fuel: 33583},
	}

	for _, tc := range tests {
		got := CalculateFuelByMass(tc.mass)
		if got != tc.fuel {
			t.Fatalf("Expected: %d, Got: %d\n", tc.fuel, got)
		}
	}
}
