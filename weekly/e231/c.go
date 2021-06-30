package e231

/*
func countRestrictedPaths(n int, edges [][]int) int {
	set := make([][]int, n)
	minSet := make([][]int, n)
	for i := 0; i < n; i++ {
		set[i] = make([]int, n)
		minSet[i] = make([]int, n)
	}

	for _, line := range edges {
		larger := -1
		smaller := -1
		weight := line[2]
		if line[0] > line[1] {
			larger += line[0]
			smaller += line[1]
		} else {
			larger += line[1]
			smaller += line[0]
		}
		set[smaller][larger] = weight
	}

}

func getMinSet(set, minSet [][]int, from, to int) {
	if minSet[from][to] > 0 {

	}
}

func wayCount(set [][]int, current, lastWeigth int) int {
	for smallerNode, weight := range set[current] {
		if weight == 0 || weight < lastWeigth {

		}
	}
}
*/
