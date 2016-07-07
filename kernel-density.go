package maths

import (
	"fmt"
	"math"
)

type Percentage float64

func (p Percentage) String() string {
	return fmt.Sprintf("%f%%", float64(p*100))
}

func (p Percentage) Float64() float64 {
	return float64(p)
}

type Dataset []float64

func NewDataset(d []float64) Dataset {
	return Dataset(d)
}

func (d Dataset) Sum() float64 {
	var sum float64
	for _, point := range d {
		sum += point
	}

	return sum
}

func (d Dataset) Avg() float64 {
	var sum float64

	for _, point := range d {
		sum += point
	}

	return sum / float64(len(d))
}

func (d Dataset) Max() float64 {
	var result float64
	for _, point := range d {
		if point > result {
			result = point
		}
	}

	return result
}

func (d Dataset) Min() float64 {
	result := d[0]
	for _, point := range d {
		if point < result {
			result = point
		}
	}

	return result
}

func (d Dataset) KernelDensity(bandwidth float64, resolution float64, kernelFunc KernelFunc) KernelDensityResult {
	max := d.Max()
	step := max / resolution

	var results Dataset
	for i := float64(0); i < max; i += step {
		var sum float64
		for _, point := range d {
			sum += kernelFunc((i - point) / bandwidth)
		}

		results = append(results, ((1 / (float64(len(d)) * bandwidth)) * sum))
	}

	return KernelDensityResult{d, results, bandwidth, resolution, kernelFunc}
}

type KernelDensityResult struct {
	dataset    Dataset
	results    Dataset
	bandwidth  float64
	resolution float64
	kernelFunc KernelFunc
}

func (d *KernelDensityResult) Results() Dataset {
	return d.dataset
}

func (d *KernelDensityResult) Plot() {
	fmt.Println("Potting not YET supported")
}

func (d KernelDensityResult) OutlierProbability(newPoint float64) Percentage {
	var sum float64
	for _, point := range d.dataset {
		sum += d.kernelFunc((newPoint - point) / d.bandwidth)
	}
	sum = (1 / (float64(len(d.dataset)) * d.bandwidth)) * sum
	resultSum := d.results.Sum()
	sum = sum / resultSum

	var normalizedResults Dataset
	for _, point := range d.results {
		normalizedResults = append(normalizedResults, point/resultSum)
	}

	var probability float64
	for _, point := range normalizedResults {
		if point > sum {
			probability += point
		}
	}

	return Percentage(probability)
}

type KernelFunc func(float64) float64

var Gaussian = func(x float64) float64 {
	return (1 / math.Sqrt(2*math.Pi)) * math.Exp(-0.5*(x*x))
}

/* NOTES
   https://en.wikipedia.org/wiki/Kernel_(statistics)#In_non-parametric_statistics
*/
