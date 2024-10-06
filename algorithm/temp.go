package algorithm

func constructGridLayout(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	degree := make(map[int][]int)
	for i, v := range g {
		degree[len(v)] = append(degree[len(v)], i)
	}

	minKey, maxKey := -1, -1
	for key := range degree {
		if minKey == -1 || key < minKey {
			minKey = key
		}
		if maxKey == -1 || key > maxKey {
			maxKey = key
		}
	}

	row := []int{}
	if minKey == 1 {
		row = append(row, degree[1][0])
	} else if maxKey <= 3 {
		// 矩阵只有两列
		x := degree[2][0]
		for _, y := range g[x] {
			if len(g[y]) == 2 {
				row = append(row, x)
				row = append(row, y)
				break
			}
		}
	} else {
		// 矩形至少有三列
		x := degree[2][0]
		row = append(row, x)
		pre := x
		x = g[x][0]
		for len(g[x]) > 2 {
			row = append(row, x)
			for _, y := range g[x] {
				if y != pre && len(g[y]) < 4 {
					pre = x
					x = y
					break
				}
			}
		}
		row = append(row, x)
	}
	ans := [][]int{row}
	vis := make([]bool, n)
	for i := range vis {
		vis[i] = false
	}
	count := n/len(row) - 1
	for i := 0; i < count; i++ {
		for _, x := range row {
			vis[x] = true
		}
		nxt_row := []int{}
		for _, x := range row {
			for _, y := range g[x] {
				if !vis[y] {
					nxt_row = append(nxt_row, y)
					break
				}
			}
		}
		ans = append(ans, nxt_row)
		row = nxt_row
	}
	return ans
}

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
