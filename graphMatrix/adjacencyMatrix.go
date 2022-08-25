package main

import "fmt"

// Graph
// type Graph struct {
// 	matrix  [][]int
// 	visited [4][5]bool
// }

// CheckRiverSize
func CheckRiverSize(matrix [][]int) []int {
	// visited := [len(matrix)][len(matrix[0])]bool{}
	visited := [4][5]bool {}
	sizes := []int{}
	visited[0][0] = false
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if visited [i][j] {
					fmt.Println(visited[i][j])
					continue
				}
				sizes = append(sizes, TraverseThroughRiver(i,j,sizes)...)
				// sizes = g.TraverseThroughRiver()
			}
		}
	return sizes
}

// TraverseThroughRiver
func TraverseThroughRiver(i int, j int, sizes []int) []int  {
	currentRiverSize := 0
	stack := [][]int{}
	stack = append(stack, []int{i, j})
	fmt.Println(stack)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i = top[0]
		j = top[1]
		// fmt.Println(i, j)
		if visited[i][j] {
			fmt.Println(visited[i][j])
			fmt.Println(visited[i][j])
			continue
		}
		CheckVisitedOrNot(i, j)
		// visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}
		fmt.Println(i, j)
		currentRiverSize++
		nearest := GetNearestOfNodes(i, j)
		for _, v := range nearest {
			stack = append(stack, v)
		}
		fmt.Println("  ",stack)
	}
	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}
	return sizes
}
func CheckVisitedOrNot(i,j int) {
	
	visited[i][j] = true
}

// GetNearestOfNodes
func GetNearestOfNodes(i, j int) [][]int {
	nearest := [][]int{}
	if i > 0 && !visited[i-1][j] {
		nearest = append(nearest, []int{i-1, j})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		nearest = append(nearest, []int{i+1, j})
	}
	if j > 0 && !visited[i][j-1] {
		nearest = append(nearest, []int{i, j-1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		nearest = append(nearest, []int{i, j+1})
	}
	return nearest
}

func main() {

	matrix := [][]int{{1, 0, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1}}
	
	result := CheckRiverSize(matrix)
	// fmt.Println(result)
	fmt.Println(result[1])
}
