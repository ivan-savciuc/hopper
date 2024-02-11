package main

// Input

import (
	"fmt"
	"strings"

	hopper "github.com/ivan-savciuc/hopper/internal"
)

func main() {
	// Define testcases as a string
	testCases := `3
	5 5
	4 0 4 4
	1
	1 4 2 3
	3 3
	0 0 2 2
	1 1 0 2
	0 2 1 1
	30 30
	0 0 29 29
	0
`

	var numTestCases int
	fmt.Sscan(testCases, &numTestCases)
	testCases = testCases[strings.Index(testCases, "\n")+1:]

	fmt.Println("Number of test cases:", numTestCases)

	for i := 0; i < numTestCases; i++ {
		var width, height, startX, startY, finishX, finishY, numObstacles int
		fmt.Sscan(testCases, &width, &height, &startX, &startY, &finishX, &finishY, &numObstacles)

		fmt.Println("Test case", i+1)
		fmt.Println("Width:", width)
		fmt.Println("Height:", height)
		fmt.Println("Start:", startX, startY)
		fmt.Println("Finish:", finishX, finishY)
		fmt.Println("Number of obstacles:", numObstacles)

		obstacles := make([]hopper.Point, numObstacles)
		for j := 0; j < numObstacles; j++ {
			fmt.Sscan(testCases, &obstacles[j].X, &obstacles[j].Y)
		}

		start := hopper.Point{X: startX, Y: startY}
		finish := hopper.Point{X: finishX, Y: finishY}

		hopper.PrintGrid(height, width, start, finish, obstacles)
		result := hopper.BFS(start, finish, obstacles, width, height)
		if result == -1 {
			fmt.Println("No solution.")
		} else {
			fmt.Printf("Optimal solution takes %d hops.\n", result)
		}

		// Remove 4 lines from testCases
		testCases = testCases[strings.Index(testCases, "\n")+1:]
		testCases = testCases[strings.Index(testCases, "\n")+1:]
		testCases = testCases[strings.Index(testCases, "\n")+1:]
		testCases = testCases[strings.Index(testCases, "\n")+1:]
	}
}
