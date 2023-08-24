package mymath

import "testing"

func TestSqrt(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want int
	}{
		// TODO: Add test cases.
		{"1", 4, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		// TODO: Add test cases.
		{"1", 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcos(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		// TODO: Add test cases.
		{"1", 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acos(tt.x); got != tt.want {
				t.Errorf("Acos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcosh(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		// TODO: Add test cases.
		{"1", 4, 2.0634370688955608},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Acosh(tt.x); got != tt.want {
				t.Errorf("Acosh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsin(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		// TODO: Add test cases.
		{"1", 1, 1.5707963267948966},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Asin(tt.x); got != tt.want {
				t.Errorf("Asin() = %v, want %v", got, tt.want)
			}
		})
	}
}
