package maths

import "math"

func KernelDensity(bandwidth float64, resolution float64, kernelFunc KernelFunc, dataset []float64) []float64 {
	largestInDataSet := GetGreatest(dataset)

	var results []float64
	rng := largestInDataSet / resolution
	for i := float64(0); i < largestInDataSet; i += rng {
		var sum float64
		for _, point := range dataset {
			sum += 1 / math.Sqrt(2*math.Pi) * math.Exp(-0.5*math.Pow(2, ((i-point)/bandwidth)))
		}

		result := kernelFunc(bandwidth, resolution, dataset, sum)
		results = append(results, result)
	}

	return results
}

type KernelFunc func(bandwidth float64, resolution float64, dataset []float64, values ...float64) float64

var Gaussian = func(bandwidth float64, resolution float64, dataset []float64, values ...float64) float64 {
	return 1 / (float64(len(dataset)) * bandwidth) * values[0]
}

/* NOTES
   https://en.wikipedia.org/wiki/Kernel_(statistics)#In_non-parametric_statistics
*/
