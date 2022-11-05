package main

import "fmt"

func main() {
	test := &Graph{}

	for i := 1; i <= 15; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(2, 4)
	test.AddEdge(2, 5)
	test.AddEdge(3, 6)
	test.AddEdge(3, 7)
	test.AddEdge(4, 8)
	test.AddEdge(4, 9)
	test.AddEdge(5, 10)
	test.AddEdge(5, 11)
	test.AddEdge(6, 12)
	test.AddEdge(6, 13)
	test.AddEdge(7, 14)
	test.AddEdge(7, 15)

	test.Print()

	fmt.Println(test.BFS(1, 8))
	fmt.Println(test.DFS(1, 11))
}

func (g *Graph) BFS(start, search any) bool {
	open := []any{start}
	closed := []any{}

	for open != nil {
		// debug open and closed
		fmt.Printf("Open: %v - Closed: %v \n", open, closed)

		current := open[0]

		if len(open) > 1 {
			open = open[1:]
		} else {
			open = nil
		}

		if current == search {
			fmt.Printf("%v : Found!\n", search)
			return true
		} else {
			for _, v := range g.vertices {
				if v.key == current {
					for _, v := range v.adjacent {
						open = append(open, v.key)
					}
				}
			}
			closed = append(closed, current)
		}
	}

	fmt.Printf("%v : Not Found!\n", search)
	return false
}

func (g *Graph) DFS(start, search any) bool {
	open := []any{start}
	closed := []any{}

	for open != nil {
		// debug open and closed
		fmt.Printf("Open: %v - Closed: %v \n", open, closed)

		current := open[0]

		if len(open) > 1 {
			open = open[1:]
		} else {
			open = nil
		}

		if current == search {
			fmt.Printf("%v : Found!\n", search)
			return true
		} else {
			for _, v := range g.vertices {
				if v.key == current {
					temp := []any{}
					for _, v := range v.adjacent {
						temp = append(temp, v.key)
					}
					open = append(temp, open...)
				}
			}
			closed = append(closed, current)
		}

	}

	fmt.Printf("%v : Not Found!\n", search)
	return false
}

// graph structure
type Graph struct {
	vertices []*Vertex
}

// vertex structure
type Vertex struct {
	key      any
	adjacent []*Vertex
}

// add vertex
func (g *Graph) AddVertex(k any) {
	if exist(g.vertices, k) {
		err := fmt.Errorf("Vertex %v already exist", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// add edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if exist(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// get vertex
func (g *Graph) getVertex(k any) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// exist
func exist(s []*Vertex, k any) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// print graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}
