package math

var mem = make(map[float64]float64)

func Average(xs []float64) float64 {

	if len(xs) == 0 {
		return 0
	}

	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

func Min(xs []float64) float64 {

	if len(xs) == 0 {
		return 0
	}

	minValue := xs[0]

	for _, x := range xs {
		if x < minValue {
			minValue = x
		}
	}

	return minValue
}

func Max(xs []float64) float64 {

	if len(xs) == 0 {
		return 0
	}
	
	maxValue := xs[0]

	for _, x := range xs {
		if x > maxValue {
			maxValue = x
		}
	}

	return maxValue
}

func Fib(n float64) float64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if value, ok := mem[n]; ok {
		return value
	}

	mem[n] = Fib(n-1) + Fib(n-2)
	return mem[n]
}
