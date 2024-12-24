package main

import (
	"github.com/gofiber/fiber/v2"
)

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PathRequest struct {
	Start Coordinates `json:"start"`
	End   Coordinates `json:"end"`
}

func main() {
	app := fiber.New()

	// API endpoint to find the shortest path
	app.Post("/find-path", func(c *fiber.Ctx) error {
		var request PathRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		var Grid [][]int

		// create a 20 x 20 grid
		for i := 0; i < 20; i++ {
			row := make([]int, 20)
			Grid = append(Grid, row)
		}

		path := findShortestPathDFS(Grid, request.Start, request.End)

		if path == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "No path found",
			})
		}

		return c.JSON(fiber.Map{
			"path": path,
		})
	})

	app.Listen(":5000")
}

func findShortestPathDFS(grid [][]int, start, end Coordinates) []Coordinates {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var result []Coordinates
	var shortestPath []Coordinates

	var dfs func(x, y int, path []Coordinates)
	dfs = func(x, y int, path []Coordinates) {
		// Base cases
		if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] == 1 || visited[x][y] {
			return
		}
		if x == end.X && y == end.Y {
			path = append(path, Coordinates{X: x, Y: y})
			if len(shortestPath) == 0 || len(path) < len(shortestPath) {
				shortestPath = make([]Coordinates, len(path))
				copy(shortestPath, path)
			}
			return
		}

		// Mark the current cell as visited
		visited[x][y] = true
		path = append(path, Coordinates{X: x, Y: y})

		// Explore all 4 directions
		dfs(x+1, y, path)
		dfs(x-1, y, path)
		dfs(x, y+1, path)
		dfs(x, y-1, path)

		// Backtrack
		visited[x][y] = false
		path = path[:len(path)-1]
	}

	dfs(start.X, start.Y, result)
	return shortestPath
}
