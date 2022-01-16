package abtest

import (
	"math"
	"testing"
)

func TestZscore(t *testing.T) {
	want := 2.6534
	got := zScore(80000, 80000, 1600, 1752)
	if math.Abs(want-got) > 0.0005 {
		t.Errorf("Zscore(%v,%v,%v,%v) -> %v, want %v", 80000, 80000, 1600, 1752, got, want)
	}
}

func TestStatisticallySignificant(t *testing.T) {

	tests := []struct {
		confidence     float64
		oA, oB, cA, cB int
		want           bool
	}{
		{0.95, 10000, 20000, 1500, 3200, true},
		{0.99, 10000, 20000, 1500, 3200, false},
	}

	for _, test := range tests {
		got := StatisticallySignificant(test.oA, test.cA, test.oB, test.cB, test.confidence)
		if got != test.want {
			t.Errorf("StatisticallySignificant(%v, %v, %v, %v, %v) -> %v, want %v", test.confidence, test.oA, test.oB, test.cA, test.cB, got, test.want)
		}
	}
}
