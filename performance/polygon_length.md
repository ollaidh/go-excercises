# Compute polygon length

There's a simple Go program that creates a Polygon and computes its length.

On a MacBook Air M2 this code runs for about ~350ms. Your task is to speed up
the length computation

Don't worry about parallelizing the code, more than 10x speedup is achievable 
without any parallelization here. The flow of the program should stay untouched
(although it might be confusing from the sanity standpoint, the flow gives you
some pointers on where to look for a bottleneck):

- Create a Polygon
- Shuffle the Polygon
- Start a timer
- Loop over all the points and accumulate the length
- Print the elapsed time


```go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X, Y float64
}

func distance(p1, p2 *Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

type Polygon struct {
	Vertices []*Point
}

func (p *Polygon) Length() float64 {
	if len(p.Vertices) < 2 {
		return 0
	}

	var totalLength float64
	for i := 0; i < len(p.Vertices)-1; i++ {
		totalLength += distance(p.Vertices[i], p.Vertices[i+1])
	}
	return totalLength
}

func makePolygon(n int) Polygon {
	vertices := make([]*Point, n)
	for i := 0; i < n; i++ {
		vertices[i] = &Point{X: rand.Float64(), Y: rand.Float64()}
	}
	rand.Shuffle(len(vertices), func(i, j int) {
		vertices[i], vertices[j] = vertices[j], vertices[i]
	})
	return Polygon{Vertices: vertices}
}

func main() {
	polygon := makePolygon(50_000_000)
	start := time.Now()
	_ = polygon.Length()
	fmt.Printf("Elapsed: %.1fms\n", time.Since(start).Seconds()*1000)
}
```