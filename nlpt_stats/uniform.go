package nlpt_stats

import "math/rand"

var NextUniform func() float64 = rand.Float64

func Uniform() func() float64 { return NextUniform }
