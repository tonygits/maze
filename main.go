package main

import "fmt"

func main() {
	fmt.Println(countWays(3, 3))
	fmt.Println(countWithDiagonal(3, 3))
	paths("", 3, 3)
	fmt.Println(pathArray("", 3, 3))
	fmt.Println(pathArrayDiagonal("", 3, 3))

	board := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}
	fmt.Println(pathWithObstacles("", board, 0, 0))
	fmt.Println(allDirectionsWithObstacles("", board, 0, 0))
}

func countWays(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}

	left := countWays(r-1, c)
	right := countWays(r, c-1)
	return left + right
}

func countWithDiagonal(r, c int) int {
	if r == 1 || c == 1 {
		return 1
	}

	diag := 0
	if r > 0 && c > 0 {
		diag = countWays(r-1, c-1)
	}
	left := countWays(r-1, c)
	right := countWays(r, c-1)
	return left + right + diag
}

func paths(p string, r, c int) {
	if r == 1 && c == 1 {
		fmt.Println(p)
		return
	}

	if r > 1 {
		paths(p+"D", r-1, c)
	}

	if c > 1 {
		paths(p+"R", r, c-1)
	}

}

func pathArray(p string, r, c int) []string {
	if r == 1 && c == 1 {
		mypaths := make([]string, 0)
		mypaths = append(mypaths, p)
		return mypaths
	}

	mypaths := make([]string, 0)
	if r > 1 {
		mypaths = append(mypaths, pathArray(p+"V", r-1, c)...)
	}

	if c > 1 {
		mypaths = append(mypaths, pathArray(p+"R", r, c-1)...)
	}

	return mypaths
}

func pathArrayDiagonal(p string, r, c int) []string {
	if r == 1 && c == 1 {
		mypaths := make([]string, 0)
		mypaths = append(mypaths, p)
		return mypaths
	}

	mypaths := make([]string, 0)
	if r > 1 && c > 1 {
		mypaths = append(mypaths, pathArrayDiagonal(p+"D", r-1, c-1)...)
	}

	if r > 1 {
		mypaths = append(mypaths, pathArrayDiagonal(p+"V", r-1, c)...)
	}

	if c > 1 {
		mypaths = append(mypaths, pathArrayDiagonal(p+"R", r, c-1)...)
	}

	return mypaths
}

func pathWithObstacles(p string, maze [][]bool, r, c int) []string {
	if r == len(maze)-1 && c == len(maze[0])-1 {
		mypaths := make([]string, 0)
		mypaths = append(mypaths, p)
		return mypaths
	}

	if !maze[r][c] {
		return []string{}
	}

	mypaths := make([]string, 0)
	if r < len(maze)-1 {
		mypaths = append(mypaths, pathWithObstacles(p+"V", maze, r+1, c)...)
	}

	if c < len(maze[0])-1 {
		mypaths = append(mypaths, pathWithObstacles(p+"R", maze, r, c+1)...)
	}

	return mypaths
}

func allDirectionsWithObstacles(p string, maze [][]bool, r, c int) []string {
	if r == len(maze)-1 && c == len(maze[0])-1 {
		mypaths := make([]string, 0)
		mypaths = append(mypaths, p)
		return mypaths
	}

	if !maze[r][c] {
		return []string{}
	}

	maze[r][c] = false

	mypaths := make([]string, 0)
	if r < len(maze)-1 {
		mypaths = append(mypaths, allDirectionsWithObstacles(p+"D", maze, r+1, c)...)
	}

	if c < len(maze[0])-1 {
		mypaths = append(mypaths, allDirectionsWithObstacles(p+"R", maze, r, c+1)...)
	}

	if r > 0 {
		mypaths = append(mypaths, allDirectionsWithObstacles(p+"U", maze, r-1, c)...)
	}

	if c > 0 {
		mypaths = append(mypaths, allDirectionsWithObstacles(p+"L", maze, r, c-1)...)
	}

	maze[r][c] = true

	return mypaths
}
