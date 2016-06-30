package maths

func GetGreatest(dataset []float64) float64 {
	var result float64
	for _, point := range dataset {
		if point > result {
			result = point
		}
	}

	return result
}
