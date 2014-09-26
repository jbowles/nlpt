package nlp_stats

import (
	"fmt"
	"testing"
)

// Randomness comes in each definition of the Multinomial function.
// Once defined, the results can be predicted and so we can test it.
func TestMultinomialDiceThrowOf20(t *testing.T) {
	//throw a dice 20 times
	dice := []float64{}

	for i := 0; i < 6; i++ {
		dice = append(dice, 1.0/6.0)
	}

	container := make(map[int][]int64)

	res := Multinomial(dice, 20)
	for i := 0; i <= 6; i++ {
		container[i] = res()
	}

	//map[0:[3 6 4 3 3 1] 1:[4 5 1 4 3 3] 2:[2 8 2 3 2 3] 3:[1 1 3 3 6 6] 4:[3 5 3 4 3 2] 5:[4 0 4 6 4 2] 6:[3 4 2 4 5 2]]
	fmt.Printf("%v\n", container)

	if container[0][5] != int64(1) {
		t.Log("first definition 0index should have been [3 6 4 3 3 1], but got slice:", container[0])
	}

	if container[1][4] != int64(3) {
		t.Log("first definition 1index should have been [4 5 1 4 3 3], but got slice:", container[1])
	}

	if container[2][3] != int64(3) {
		t.Log("first definition 2index should have been [2 8 2 3 2 3], but got slice:", container[2])
	}

	if container[3][2] != int64(3) {
		t.Log("first definition 3index should have been [1 1 3 3 6 6], but got slice:", container[3])
	}

	if container[4][1] != int64(5) {
		t.Log("first definition 4index should have been [3 5 3 4 3 2], but got slice:", container[4])
	}

	if container[5][0] != int64(4) {
		t.Log("first definition 5index should have been [4 0 4 6 4 2], but got slice:", container[5])
	}

	if container[6][5] != int64(2) {
		t.Log("first definition 6index should have been [3 4 2 4 5 2], but got slice:", container[6])
	}
}
