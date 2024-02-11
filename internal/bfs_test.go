package internal

import (
	"testing"
)

type TestCase struct {
	h, w      int
	start     Point
	finish    Point
	obstacles []Point
	result    int
}

func TestBFS(t *testing.T) {
	testCases := []TestCase{
		{
			5, 5,
			Point{0, 0},
			Point{4, 4},
			[]Point{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {3, 1}, {3, 2}, {3, 3}, {3, 4}},
			5,
		},
		// {
		// 	5, 5,
		// 	Point{0, 0},
		// 	Point{4, 4},
		// 	[]Point{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
		// 	-1,
		// },
		// {
		// 	5, 5,
		// 	Point{0, 0},
		// 	Point{4, 4},
		// 	[]Point{},
		// 	3,
		// },
		// {
		// 	5, 5,
		// 	Point{4, 0},
		// 	Point{0, 0},
		// 	[]Point{},
		// 	3,
		// },
		// {
		// 	5, 5,
		// 	Point{0, 0},
		// 	Point{4, 4},
		// 	[]Point{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}},
		// 	5,
		// },
		{
			30, 30,
			Point{0, 0},
			Point{24, 24},
			[]Point{},
			6,
		},
	}

	for _, testCase := range testCases {
		width, height := testCase.h, testCase.w
		start, finish := testCase.start, testCase.finish
		obstacles := testCase.obstacles

		PrintGrid(testCase.h, testCase.w, start, finish, obstacles)
		result := BFS(start, finish, obstacles, width, height)
		if result != testCase.result {
			t.Errorf("Expected %d, but got %d", testCase.result, result)
		}
	}
}
