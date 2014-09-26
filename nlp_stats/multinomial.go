package nlp_stats

var iZero int64 = int64(0)

func NextMultinomial(θ []float64, n int64) []int64 {
	x := make([]int64, len(θ))
	chooser := Choice(θ)
	for i := iZero; i < n; i++ {
		x[chooser()]++
	}
	return x
}
func Multinomial(θ []float64, n int64) func() []int64 {
	return func() []int64 {
		return NextMultinomial(θ, n)
	}
}

func NextChoice(θ []float64) int64 {
	u := NextUniform()
	i := 0
	sum := θ[0]
	for ; sum < u && i < len(θ)-1; i++ {
		sum += θ[i+1]
	}
	if u >= sum {
		return int64(len(θ))
	}
	return int64(i)
}
func Choice(θ []float64) func() int64 {
	return func() int64 {
		return NextChoice(θ)
	}
}
