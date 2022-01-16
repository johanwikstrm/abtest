// Package abtest implements a statistical significance test for AB testing.
// It has no dependencies outside the standard library.
// It uses the same statistical significance test as popular AB testing tools like https://abtestguide.com/calc/
package abtest

import (
	"math"
)

func zScore(nA, nB, convA, convB int) float64 {
	ratioA := float64(convA) / float64(nA)
	ratioB := float64(convB) / float64(nB)

	// Variance
	varControl := ratioA * (1 - ratioA)
	varVariant := ratioB * (1 - ratioB)

	sMean := ratioB - ratioA
	sVar := (varControl / float64(nA)) + (varVariant / float64(nB))
	zScore := sMean / math.Sqrt(sVar)
	return zScore
}

func normCDF(x float64) float64 {
	return 0.5 * (1 + math.Erf((x)/(math.Sqrt(2))))
}

// StatisticallySignificant returns true if B (the experiment) converted at a higher rate than A (the control) at a given confidence level.
// It uses a one-sided significance test which is the most appropriate type for AB-tests.
// https://blog.analytics-toolkit.com/2017/one-tailed-two-tailed-tests-significance-ab-testing/
func StatisticallySignificant(numObservationsA, numConvertedA, numObservationsB, numConvertedB int, confidenceLevel float64) bool {
	zScore := zScore(numObservationsA, numObservationsB, numConvertedA, numConvertedB)
	return confidenceLevel < normCDF(zScore)
}
