package main

func main() {

}

// 题目忘了
type Task struct {
	dependency []string
	this       string
}

/*
func search(tasks []Task) {
	levels := make([]map[string]struct{}, 0)

	for level := 0; len(tasks) > 0; level++ {
		levels = append(levels, make(map[string]struct{}, 0))
		for i, t := range tasks {
			if len(t.dependency) == 0 {
				levels[level] = append(levels[level], t.this)
				tasks = append(tasks[:i], tasks[i+1:]...)
				continue
			}

			for _, base := range t.dependency {
				tasks[level-1][base]
			}

		}
	}

}


*/
