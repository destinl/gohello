package algorithm

func remainingMethods(n int, k int, invocations [][]int) []int {

	edges := make(map[int][]int)
	for _, v := range invocations {
		edges[v[0]] = append(edges[v[0]], v[1])
	}

	visited := make(map[int]bool)
	queue := []int{k}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true
		for _, v := range edges[u] {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if !visited[i] {
			for _, j := range edges[i] {
				if visited[j] {
					res1 := make([]int, n)
					for i := 0; i < n; i++ {
						res1[i] = i
					}
					return res1
				}
			}
			res = append(res, i)
		}
	}
	return res
}
