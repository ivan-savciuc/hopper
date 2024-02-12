# Hopper

Hopper is a Go program that solves a specific pathfinding problem. It uses a breadth-first search (BFS) algorithm to find the shortest path from a start point to a finish point on a grid while avoiding obstacles.

## Problem Description

The hopper starts at a specific point on a grid with a 0,0 velocity. It can change its velocity by a 1 in either direction, and it could be -1/0/1 at each step. The goal is to find the shortest path to the finish point.

## Run

To run the program, you need to have Go installed on your machine. Then, you can clone this repository and build the program:

```bash
git clone https://github.com/yourusername/hopper.git
cd hopper
go build
./hopper
```

or you can use `go run main.go`

## Tests

To run the tests for the Hopper program, navigate to the directory containing the source code and execute the following command:

```bash
time go test ./... -v  
```

### Thoughts about the implementation

I was moving from one place to another last Thursday and Friday and had no chance to ask questions since I started working on this assignment over the weekend. So I hope I understood it right :)

The first problem with BFS implementations is that it takes significant time on an empty grid, like 30x30, where the start and finish points are on opposite sides of the greed.

I tried several approaches to make it faster for such use cases, and one was to run more things in parallel; however, every iteration in the loop is pretty trivial, and concurrency overhead wasn't worth it.

My second idea was to direct velocity towards the finish point and employ a priority queue for such cases. It makes the BFS implementation significantly faster, but some tests show wrong results on 1-2 hopps more than an optimal route.

All in all, I'm submitting a fairly simple BFS implementation due to time restrictions, and I had lots of fun working on this. Thank you :)
