package main

func main() {

}

type Assist struct {
	graph [][]int
}

// 魔塔问题，no ac
/*
{-5,-1,2},
{-10,-3,-1},
{0,30,-2},
*/
func (a *Assist) move(i, j, current, min int) (res int) {
	current += a.graph[i][j]
	if a.graph[i][j] < min {
		min = a.graph[i][j]
	}

	resI, resJ := 0, 0
	if i < len(a.graph) {
		resI = a.move(i+1, j, current, min)
	}

	if j < len(a.graph[i]) {
		resJ = a.move(i, j+1, current, min)
	}

	if current < min {

	}

	if resI > resJ {
		res = resI + a.graph[i][j]
		return
	}
	res = resI + a.graph[i][j]
	return
}
